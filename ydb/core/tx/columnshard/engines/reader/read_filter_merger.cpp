#include "read_filter_merger.h"
#include <library/cpp/actors/core/log.h>

namespace NKikimr::NOlap::NIndexedReader {

NJson::TJsonValue TSortableBatchPosition::DebugJson() const {
    NJson::TJsonValue result;
    result["reverse"] = ReverseSort;
    result["records_count"] = RecordsCount;
    result["position"] = Position;
    result["sorting"] = Sorting->DebugJson(Position);
    if (Data) {
        result["data"] = Data->DebugJson(Position);
    }
    return result;
}

std::optional<ui64> TSortableBatchPosition::FindPosition(std::shared_ptr<arrow::RecordBatch> batch, const TSortableBatchPosition& forFound, const bool greater, const bool include) {
    if (!batch || !batch->num_rows()) {
        return {};
    }

    const auto checkEqualBorder = [batch, greater, include](const i64 position) ->std::optional<i64> {
        if (include) {
            return position;
        } else if (greater) {
            if (batch->num_rows() > position + 1) {
                return position + 1;
            } else {
                return {};
            }
        } else {
            if (position) {
                return position - 1;
            } else {
                return {};
            }
        }
    };

    i64 posStart = 0;
    i64 posFinish = batch->num_rows() - 1;
    TSortableBatchPosition position = forFound.BuildSame(batch, posStart);
    {
        position.InitPosition(posStart);
        auto cmp = position.Compare(forFound);
        if (cmp == std::partial_ordering::greater) {
            if (greater) {
                return posStart;
            } else {
                return {};
            }
        } else if (cmp == std::partial_ordering::equivalent) {
            return checkEqualBorder(posStart);
        }
    }
    {
        position.InitPosition(posFinish);
        auto cmp = position.Compare(forFound);
        if (cmp == std::partial_ordering::less) {
            if (greater) {
                return {};
            } else {
                return posFinish;
            }
        } else if (cmp == std::partial_ordering::equivalent) {
            return checkEqualBorder(posFinish);
        }
    }
    while (posFinish > posStart + 1) {
        Y_VERIFY(position.InitPosition(0.5 * (posStart + posFinish)));
        const auto comparision = position.Compare(forFound);
        if (comparision == std::partial_ordering::less) {
            posStart = position.Position;
        } else if (comparision == std::partial_ordering::greater) {
            posFinish = position.Position;
        } else {
            return checkEqualBorder(position.Position);
        }
    }
    Y_VERIFY(posFinish != posStart);
    if (greater) {
        return posFinish;
    } else {
        return posStart;
    }
}

void TMergePartialStream::PutControlPoint(std::shared_ptr<TSortableBatchPosition> point) {
    Y_VERIFY(point);
    Y_VERIFY(point->IsSameSortingSchema(SortSchema));
    Y_VERIFY(point->IsReverseSort() == Reverse);
    Y_VERIFY(++ControlPoints == 1);

    SortHeap.emplace_back(TBatchIterator(*point));
    std::push_heap(SortHeap.begin(), SortHeap.end());
}

void TMergePartialStream::AddPoolSource(const std::optional<ui32> poolId, std::shared_ptr<arrow::RecordBatch> batch, std::shared_ptr<NArrow::TColumnFilter> filter) {
    if (!batch || !batch->num_rows()) {
        return;
    }
    Y_VERIFY_DEBUG(NArrow::IsSorted(batch, SortSchema));
    if (!poolId) {
        AddNewToHeap(poolId, batch, filter, true);
    } else {
        auto it = BatchPools.find(*poolId);
        if (it == BatchPools.end()) {
            it = BatchPools.emplace(*poolId, std::deque<TIteratorData>()).first;
        }
        it->second.emplace_back(batch, filter);
        if (it->second.size() == 1) {
            AddNewToHeap(poolId, batch, filter, true);
        }
    }
}

void TMergePartialStream::AddNewToHeap(const std::optional<ui32> poolId, std::shared_ptr<arrow::RecordBatch> batch, std::shared_ptr<NArrow::TColumnFilter> filter, const bool restoreHeap) {
    if (!filter || filter->IsTotalAllowFilter()) {
        SortHeap.emplace_back(TBatchIterator(batch, nullptr, SortSchema->field_names(), DataSchema ? DataSchema->field_names() : std::vector<std::string>(), Reverse, poolId));
    } else if (filter->IsTotalDenyFilter()) {
        return;
    } else {
        SortHeap.emplace_back(TBatchIterator(batch, filter, SortSchema->field_names(), DataSchema ? DataSchema->field_names() : std::vector<std::string>(), Reverse, poolId));
    }
    if (restoreHeap) {
        std::push_heap(SortHeap.begin(), SortHeap.end());
    }
}

void TMergePartialStream::RemoveControlPoint() {
    Y_VERIFY(ControlPoints == 1);
    Y_VERIFY(ControlPointEnriched());
    Y_VERIFY(-- ControlPoints == 0);
    std::pop_heap(SortHeap.begin(), SortHeap.end());
    SortHeap.pop_back();
}

void TMergePartialStream::CheckSequenceInDebug(const TSortableBatchPosition& nextKeyColumnsPosition) {
#ifndef NDEBUG
    if (CurrentKeyColumns) {
        const bool linearExecutionCorrectness = CurrentKeyColumns->Compare(nextKeyColumnsPosition) == std::partial_ordering::less;
        if (!linearExecutionCorrectness) {
            const bool newSegmentScan = nextKeyColumnsPosition.GetPosition() == 0;
            AFL_VERIFY(newSegmentScan && nextKeyColumnsPosition.Compare(*CurrentKeyColumns) == std::partial_ordering::less)
                ("merge_debug", DebugJson())("current_ext", nextKeyColumnsPosition.DebugJson())("newSegmentScan", newSegmentScan);
        }
    }
    CurrentKeyColumns = nextKeyColumnsPosition;
#else
    Y_UNUSED(nextKeyColumnsPosition);
#endif
}

bool TMergePartialStream::DrainCurrentTo(TRecordBatchBuilder& builder, const TSortableBatchPosition& readTo, const bool includeFinish) {
    Y_VERIFY((ui32)DataSchema->num_fields() == builder.GetBuildersCount());
    PutControlPoint(std::make_shared<TSortableBatchPosition>(readTo));
    bool cpReachedFlag = false;
    while (SortHeap.size() && !cpReachedFlag) {
        if (SortHeap.front().IsControlPoint()) {
            RemoveControlPoint();
            cpReachedFlag = true;
            if (SortHeap.empty() || !includeFinish || SortHeap.front().GetKeyColumns().Compare(readTo) == std::partial_ordering::greater) {
                return true;
            }
        }

        if (auto currentPosition = DrainCurrentPosition()) {
            CheckSequenceInDebug(*currentPosition);
            builder.AddRecord(*currentPosition);
        }
    }
    return false;
}

bool TMergePartialStream::DrainAll(TRecordBatchBuilder& builder) {
    Y_VERIFY((ui32)DataSchema->num_fields() == builder.GetBuildersCount());
    while (SortHeap.size()) {
        if (auto currentPosition = DrainCurrentPosition()) {
            CheckSequenceInDebug(*currentPosition);
            builder.AddRecord(*currentPosition);
        }
    }
    return false;
}

std::optional<TSortableBatchPosition> TMergePartialStream::DrainCurrentPosition() {
    Y_VERIFY(SortHeap.size());
    Y_VERIFY(!SortHeap.front().IsControlPoint());
    TSortableBatchPosition result = SortHeap.front().GetKeyColumns();
    TSortableBatchPosition resultVersion = SortHeap.front().GetVersionColumns();
    bool isFirst = true;
    const bool deletedFlag = SortHeap.front().IsDeleted();
    while (SortHeap.size() && (isFirst || result.Compare(SortHeap.front().GetKeyColumns()) == std::partial_ordering::equivalent)) {
        auto& anotherIterator = SortHeap.front();
        if (!isFirst) {
            Y_VERIFY(resultVersion.Compare(anotherIterator.GetVersionColumns()) == std::partial_ordering::greater);
        }
        NextInHeap(true);
        isFirst = false;
    }
    if (deletedFlag) {
        return {};
    }
    return result;
}

NJson::TJsonValue TMergePartialStream::TBatchIterator::DebugJson() const {
    NJson::TJsonValue result;
    result["is_cp"] = IsControlPoint();
    if (PoolId) {
        result["pool_id"] = *PoolId;
    } else {
        result["pool_id"] = "absent";
    }
    result["key"] = KeyColumns.DebugJson();
    return result;
}

TSortableScanData::TSortableScanData(std::shared_ptr<arrow::RecordBatch> batch, const std::vector<std::string>& columns) {
    for (auto&& i : columns) {
        auto c = batch->GetColumnByName(i);
        AFL_VERIFY(c)("column_name", i)("columns", JoinSeq(",", columns));
        Columns.emplace_back(c);
        auto f = batch->schema()->GetFieldByName(i);
        Fields.emplace_back(f);
    }
}

void TRecordBatchBuilder::AddRecord(const TSortableBatchPosition& position) {
    Y_VERIFY_DEBUG(position.GetData().GetColumns().size() == Builders.size());
    Y_VERIFY_DEBUG(IsSameFieldsSequence(position.GetData().GetFields(), Fields));
//    AFL_TRACE(NKikimrServices::TX_COLUMNSHARD)("event", "record_add_on_read")("record", position.DebugJson());
    for (ui32 i = 0; i < position.GetData().GetColumns().size(); ++i) {
        NArrow::Append(*Builders[i], *position.GetData().GetColumns()[i], position.GetPosition());
    }
    ++RecordsCount;
}

}

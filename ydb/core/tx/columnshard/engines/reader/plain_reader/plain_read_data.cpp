#include "plain_read_data.h"

namespace NKikimr::NOlap::NPlainReader {

TPlainReadData::TPlainReadData(TReadMetadata::TConstPtr readMetadata, const TReadContext& context)
    : TBase(context, readMetadata)
    , EFColumns(std::make_shared<TColumnsSet>(GetReadMetadata()->GetEarlyFilterColumnIds(), GetReadMetadata()->GetIndexInfo()))
    , PKColumns(std::make_shared<TColumnsSet>(GetReadMetadata()->GetPKColumnIds(), GetReadMetadata()->GetIndexInfo()))
    , FFColumns(std::make_shared<TColumnsSet>(GetReadMetadata()->GetAllColumns(), GetReadMetadata()->GetIndexInfo()))
    , TrivialEFFlag(EFColumns->ColumnsOnly(GetReadMetadata()->GetIndexInfo().ArrowSchemaSnapshot()->field_names()))
{
    PKFFColumns = std::make_shared<TColumnsSet>(*PKColumns + *FFColumns);
    EFPKColumns = std::make_shared<TColumnsSet>(*EFColumns + *PKColumns);
    FFMinusEFColumns = std::make_shared<TColumnsSet>(*FFColumns - *EFColumns);
    FFMinusEFPKColumns = std::make_shared<TColumnsSet>(*FFColumns - *EFColumns - *PKColumns);

    Y_VERIFY(FFColumns->Contains(EFColumns));
    ui32 sourceIdx = 0;
    std::deque<std::shared_ptr<IDataSource>> sources;
    const auto& portionsOrdered = GetReadMetadata()->SelectInfo->GetPortionsOrdered(GetReadMetadata()->IsDescSorted());
    const auto& committed = readMetadata->CommittedBlobs;
    auto itCommitted = committed.begin();
    auto itPortion = portionsOrdered.begin();
    ui64 portionsBytes = 0;
    while (itCommitted != committed.end() || itPortion != portionsOrdered.end()) {
        bool movePortion = false;
        if (itCommitted == committed.end()) {
            movePortion = true;
        } else if (itPortion == portionsOrdered.end()) {
            movePortion = false;
        } else if (itCommitted->GetFirstVerified() < (*itPortion)->IndexKeyStart()) {
            movePortion = false;
        } else {
            movePortion = true;
        }

        if (movePortion) {
            portionsBytes += (*itPortion)->BlobsBytes();
            auto start = GetReadMetadata()->BuildSortedPosition((*itPortion)->IndexKeyStart());
            auto finish = GetReadMetadata()->BuildSortedPosition((*itPortion)->IndexKeyEnd());
            AFL_DEBUG(NKikimrServices::TX_COLUMNSHARD)("event", "portions_for_merge")("start", start.DebugJson())("finish", finish.DebugJson());
            sources.emplace_back(std::make_shared<TPortionDataSource>(sourceIdx++, (*itPortion), *this, start, finish));
            ++itPortion;
        } else {
            auto start = GetReadMetadata()->BuildSortedPosition(itCommitted->GetFirstVerified());
            auto finish = GetReadMetadata()->BuildSortedPosition(itCommitted->GetLastVerified());
            sources.emplace_back(std::make_shared<TCommittedDataSource>(sourceIdx++, *itCommitted, *this, start, finish));
            ++itCommitted;
        }
    }
    Scanner = std::make_shared<TScanHead>(std::move(sources), *this);

    auto& stats = GetReadMetadata()->ReadStats;
    stats->IndexPortions = GetReadMetadata()->SelectInfo->PortionsOrderedPK.size();
    stats->IndexBatches = GetReadMetadata()->NumIndexedBlobs();
    stats->CommittedBatches = GetReadMetadata()->CommittedBlobs.size();
    stats->SchemaColumns = GetReadMetadata()->GetSchemaColumnsCount();
    stats->PortionsBytes = portionsBytes;

}

std::vector<NKikimr::NOlap::TPartialReadResult> TPlainReadData::DoExtractReadyResults(const int64_t maxRowsInBatch) {
    Scanner->DrainResults();
    if (ReadyResultsCount < maxRowsInBatch && !Scanner->IsFinished()) {
        return {};
    }
    ReadyResultsCount = 0;

    auto result = TPartialReadResult::SplitResults(PartialResults, maxRowsInBatch);
    PartialResults.clear();
    ui32 count = 0;
    for (auto&& r: result) {
        r.BuildLastKey(ReadMetadata->GetIndexInfo().GetReplaceKey());
        r.StripColumns(ReadMetadata->GetResultSchema());
        count += r.GetRecordsCount();
        r.ApplyProgram(ReadMetadata->GetProgram());
    }
    AFL_DEBUG(NKikimrServices::TX_COLUMNSHARD)("event", "DoExtractReadyResults")("result", result.size())("count", count)("finished", Scanner->IsFinished());
    return result;
}

void TPlainReadData::DoAddData(const TBlobRange& blobRange, const TString& data) {
    AFL_TRACE(NKikimrServices::TX_COLUMNSHARD)("event", "DoAddData")("range", blobRange);
    auto it = Sources.find(blobRange);
    Y_VERIFY(it != Sources.end());
    TString dataForMove = data;
    it->second->AddData(blobRange, std::move(dataForMove));
    Sources.erase(it);
}

std::optional<NKikimr::NOlap::TBlobRange> TPlainReadData::DoExtractNextBlob(const bool /*hasReadyResults*/) {
    while (Queue.empty() && Scanner->BuildNextInterval()) {
    }
    auto blobRange = Queue.pop_front();
    if (blobRange) {
        AFL_TRACE(NKikimrServices::TX_COLUMNSHARD)("event", "DoExtractNextBlob")("range", *blobRange);
    } else {
        AFL_TRACE(NKikimrServices::TX_COLUMNSHARD)("event", "DoExtractNextBlob")("range", "nothing");
    }
    return blobRange;
}

void TPlainReadData::OnIntervalResult(std::shared_ptr<arrow::RecordBatch> batch) {
    if (batch && batch->num_rows()) {
        TPartialReadResult result(std::make_shared<TScanMemoryLimiter::TGuard>(Context.GetMemoryAccessor()), batch);
        ReadyResultsCount += result.GetRecordsCount();
        PartialResults.emplace_back(std::move(result));
    }
}

NKikimr::NOlap::NPlainReader::TFetchingPlan TPlainReadData::GetColumnsFetchingPlan(const bool exclusiveSource) const {
    if (exclusiveSource) {
        if (Context.GetIsInternalRead()) {
            return TFetchingPlan(FFColumns, EmptyColumns, true);
        } else {
            if (TrivialEFFlag) {
                return TFetchingPlan(FFColumns, EmptyColumns, true);
            } else {
                return TFetchingPlan(EFColumns, FFMinusEFColumns, true);
            }
        }
    } else {
        if (GetContext().GetIsInternalRead()) {
            return TFetchingPlan(PKFFColumns, EmptyColumns, false);
        } else {
            if (TrivialEFFlag) {
                return TFetchingPlan(PKFFColumns, EmptyColumns, false);
            } else {
                return TFetchingPlan(EFPKColumns, FFMinusEFPKColumns, false);
            }
        }
    }
}

}

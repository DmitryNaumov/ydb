#include "db_wrapper.h"
#include "insert_table/insert_table.h"

#include <ydb/core/tx/columnshard/columnshard_schema.h>

#include <library/cpp/testing/unittest/registar.h>
#include <util/string/printf.h>

namespace NKikimr {

using namespace NOlap;

namespace {

class TTestInsertTableDB : public IDbWrapper {
public:
    void Insert(const TInsertedData&) override {}
    void Commit(const TInsertedData&) override {}
    void Abort(const TInsertedData&) override {}
    void EraseInserted(const TInsertedData&) override {}
    void EraseCommitted(const TInsertedData&) override {}
    void EraseAborted(const TInsertedData&) override {}

    bool Load(TInsertTableAccessor&,
              const TInstant&) override
    {
        return true;
    }

    void WriteGranule(ui32, const IColumnEngine&, const TGranuleRecord&) override {}
    void EraseGranule(ui32, const IColumnEngine&, const TGranuleRecord&) override {}
    bool LoadGranules(ui32, const IColumnEngine&, const std::function<void(const TGranuleRecord&)>&) override { return true; }

    void WriteColumn(ui32, const TPortionInfo&, const TColumnRecord&) override {}
    void EraseColumn(ui32, const TPortionInfo&, const TColumnRecord&) override {}
    bool LoadColumns(ui32, const std::function<void(const TPortionInfo&, const TColumnChunkLoadContext&)>&) override { return true; }

    void WriteCounter(ui32, ui32, ui64) override {}
    bool LoadCounters(ui32, const std::function<void(ui32 id, ui64 value)>&) override { return true; }
};

}


Y_UNIT_TEST_SUITE(TColumnEngineTestInsertTable) {
    Y_UNIT_TEST(TestInsertCommit) {
        ui64 writeId = 0;
        ui64 tableId = 0;
        TString dedupId = "0";
        TUnifiedBlobId blobId1(2222, 1, 1, 100, 1);

        TTestInsertTableDB dbTable;
        TInsertTable insertTable;
        TSnapshot indexSnapshot(1, 1);

        // insert, not commited
        bool ok = insertTable.Insert(dbTable, TInsertedData(writeId, tableId, dedupId, blobId1, {}, indexSnapshot));
        UNIT_ASSERT(ok);

        // insert the same blobId1 again
        ok = insertTable.Insert(dbTable, TInsertedData(writeId, tableId, dedupId, blobId1, {}, indexSnapshot));
        UNIT_ASSERT(!ok);

        // insert different blodId with the same writeId and dedupId
        TUnifiedBlobId blobId2(2222, 1, 2, 100, 1);
        ok = insertTable.Insert(dbTable, TInsertedData(writeId, tableId, dedupId, blobId2, {}, indexSnapshot));
        UNIT_ASSERT(!ok);

        // read nothing
        auto blobs = insertTable.Read(tableId, TSnapshot::Zero(), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 0);
        blobs = insertTable.Read(tableId+1, TSnapshot::Zero(), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 0);

        // commit
        ui64 planStep = 100;
        ui64 txId = 42;
        insertTable.Commit(dbTable, planStep, txId, {TWriteId{writeId}}, [](ui64){ return true; });

        UNIT_ASSERT_EQUAL(insertTable.GetPathPriorities().size(), 1);
        UNIT_ASSERT_EQUAL(insertTable.GetPathPriorities().begin()->second.size(), 1);
        UNIT_ASSERT_EQUAL((*insertTable.GetPathPriorities().begin()->second.begin())->GetCommitted().size(), 1);

        // read old snapshot
        blobs = insertTable.Read(tableId, TSnapshot::Zero(), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 0);
        blobs = insertTable.Read(tableId+1, TSnapshot::Zero(), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 0);

        // read new snapshot
        blobs = insertTable.Read(tableId, TSnapshot(planStep, txId), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 1);
        blobs = insertTable.Read(tableId+1, TSnapshot::Zero(), nullptr);
        UNIT_ASSERT_EQUAL(blobs.size(), 0);
    }
}

}

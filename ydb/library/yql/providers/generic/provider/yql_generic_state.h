#pragma once

#include "yql_generic_settings.h"

#include <ydb/library/yql/core/yql_data_provider.h>
#include <ydb/library/yql/providers/generic/connector/libcpp/client.h>

namespace NKikimr::NMiniKQL {
    class IFunctionRegistry;
}

namespace NYql {
    struct TGenericState: public TThrRefBase {
        using TPtr = TIntrusivePtr<TGenericState>;

        using TTableAddress = std::pair<TString, TString>; // std::pair<clusterName, tableName>

        struct TTableMeta {
            const TStructExprType* ItemType = nullptr;
            TVector<TString> ColumnOrder;
            NYql::NConnector::NApi::TSchema Schema;
            NYql::NConnector::NApi::TDataSourceInstance DataSourceInstance;
        };

        using TGetTableResult = std::pair<std::optional<const TTableMeta*>, std::optional<TIssue>>;

        TGenericState() = delete;

        TGenericState(
            TTypeAnnotationContext* types,
            const NKikimr::NMiniKQL::IFunctionRegistry* functionRegistry,
            const std::shared_ptr<NYql::IDatabaseAsyncResolver>& databaseResolver,
            const NConnector::IClient::TPtr& genericClient,
            const TGatewaysConfig* gatewaysConfig = nullptr)
            : Types(types)
            , Configuration(MakeIntrusive<TGenericConfiguration>())
            , FunctionRegistry(functionRegistry)
            , DatabaseResolver(databaseResolver)
            , GenericClient(genericClient)
        {
            if (gatewaysConfig) {
                Configuration->Init(gatewaysConfig->GetGeneric(), databaseResolver, DatabaseAuth, types->Credentials);
            }
        }

        void AddTable(const TStringBuf& clusterName, const TStringBuf& tableName, TTableMeta&& tableMeta);
        TGetTableResult GetTable(const TStringBuf& clusterName, const TStringBuf& tableName) const;
        TGetTableResult GetTable(const TStringBuf& clusterName, const TStringBuf& tableName, const TPosition& position) const;

        TTypeAnnotationContext* Types;
        TGenericConfiguration::TPtr Configuration = MakeIntrusive<TGenericConfiguration>();
        const NKikimr::NMiniKQL::IFunctionRegistry* FunctionRegistry;

        // key - (database id, database type), value - credentials to access MDB API
        NYql::IDatabaseAsyncResolver::TDatabaseAuthMap DatabaseAuth;
        std::shared_ptr<NYql::IDatabaseAsyncResolver> DatabaseResolver;

        NConnector::IClient::TPtr GenericClient;

    private:
        THashMap<TTableAddress, TTableMeta> Tables_;
    };
}

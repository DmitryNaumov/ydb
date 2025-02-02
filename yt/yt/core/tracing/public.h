#pragma once

#include <yt/yt/core/misc/public.h>

#include <library/cpp/yt/misc/guid.h>

namespace NYT::NTracing {

////////////////////////////////////////////////////////////////////////////////

constexpr auto MemoryTagLiteral = "memory_tag";

////////////////////////////////////////////////////////////////////////////////

namespace NProto {

class TTracingExt;

} // namespace NProto

DECLARE_REFCOUNTED_CLASS(TTraceContext)

DECLARE_REFCOUNTED_CLASS(TTracingConfig)

DECLARE_REFCOUNTED_CLASS(TAllocationTags)

using TTraceId = TGuid;
constexpr TTraceId InvalidTraceId = {};

using TSpanId = ui64;
constexpr TSpanId InvalidSpanId = 0;

// Request ids come from RPC infrastructure but
// we should avoid include-dependencies here.
using TRequestId = TGuid;
constexpr TRequestId InvalidRequestId = {};

////////////////////////////////////////////////////////////////////////////////

} // namespace NYT::NTracing

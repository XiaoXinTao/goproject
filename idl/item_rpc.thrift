include "base.thrift"
include "item.thrift"
namespace go xintao.project.ecommerce

/*
对于商品的 增删改查
*/

struct OrderItemRequest {
    1: required i64 ItemGroupId,
    2: required i64 ItemId,
    3: required i32 ItemCount,
    4: required i64 PassportUid,
    255: base.Base Base,
}

struct OrderItemResponse {
    255: base.BaseResp BaseResp,
}

struct
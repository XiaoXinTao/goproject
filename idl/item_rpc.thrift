include "base.thrift"
namespace go xintao.project.item_rpc

/*
对于商品的 增删改查
*/
struct ItemInfo {
    1: i64 ItemId, // Item 标识Id
    2: string ItemName, // Item 名字
    3: i32 IsLimited, // Item 是否限量
    4: i32 ItemCount, // Item 剩余数目
}

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

struct GetItemRequest {
    1: i64 ItemGroupId,
    255: base.Base Base,
}

struct GetItemResponse {
    1: list<ItemInfo> ItemList,
    255: base.BaseResp BaseResp,
}

struct CreateItemGroupRequest {
    1: i64 ItemGroupId,
    2: i64 PassportUid,
    255: base.Base Base,
}

struct CreateItemGroupResponse {
    255: base.BaseResp BaseResp,
}

struct UpdateItemRequest {
    1: i64 ItemGroupId,
    2: i64 PassportUid,
    3: list<ItemInfo> ItemList,
    255: base.Base Base,
}

struct UpdateItemResponse {
    255: base.BaseResp BaseResp,
}

service EcommerceItemRpc {
    OrderItemResponse OrderItem(1: OrderItemRequest req),
    GetItemResponse GetItem(1: GetItemRequest req),
    CreateItemGroupResponse CreateItem(1: CreateItemGroupRequest req),
    UpdateItemResponse UpdateItem(1: UpdateItemRequest req),
}
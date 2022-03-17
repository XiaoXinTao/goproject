include "base.thrift"
namespace go xintao.project.ecommerce

/*
实现商品
对于商家：增加 修改
对于买家：购买（减库存）
ItemGroup 标识是哪个活动

*/

struct ItemInfo {
    1: i64 ItemId, // Item 标识Id
    2: string ItemName, // Item 名字
    3: i32 IsLimited, // Item 是否限量
    4: i32 ItemCount, // Item 剩余数目
}

struct GetItemRequest {
    1: i64 ItemGroupId,
}

struct GetItemResponse {
    1: i32 code,
    2: string msg,
    3: list<ItemInfo> data,
}

// 表明开始一个group
struct CreateItemGroupRequest {
    1: i64 ItemGroupId,
    2: i64 PassportUid,
}

struct CreateItemGroupResponse {
    1: i32 code,
    2: string msg,
}

/*
对group内的商品增加/删除
根据itemid判断
如果存在 则修改
否则 插入
*/
struct UpdateItemRequest {
    1: i64 ItemGroupId,
    2: i64 PassportUid,
    3: list<ItemInfo> ItemList,
}

struct UpdateItemResponse {
    1: i32 code,
    2: string msg,
}

/*
用户购买行为
 */
struct OrderItemRequest {
    1: required i64 ItemGroupId,
    2: required i64 ItemId,
    3: required i32 ItemCount,
    4: required i64 PassportUid,
}

struct OrderItemResponse {
    1: i32 code,
    2: string msg,
}

service EcommerceItem {
    GetItemResponse GetItem(1: GetItemRequest req) (api.get="/xintao/project/ecommerce/get_item"),
    CreateItemGroupResponse CreateItemGroup(1: CreateItemGroupRequest req) (api.post="/xintao/project/ecommerce/create_item_group"),
    UpdateItemResponse UpdateItem(1: UpdateItemRequest req) (api.post="/xintao/project/ecommerce/update_item")
    OrderItemResponse OrderItem(1: OrderItemRequest req) (api.post="/xintao/project/ecommerce/order_item")
}

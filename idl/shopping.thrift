include "base.thrift"
namespace go xintao.project.ecommerce

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

service EcommerceShopping {
    OrderItemResponse OrderItem(1: OrderItemRequest req) (api.post="/xintao/project/ecommerce/order_item")
}

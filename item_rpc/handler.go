package main

import (
	"context"

	"github.com/XiaoXinTao/goproject/item_rpc/handler"
	"github.com/XiaoXinTao/goproject/item_rpc/kitex_gen/xintao/project/item_rpc"
)

// EcommerceItemRpcImpl implements the last service interface defined in the IDL.
type EcommerceItemRpcImpl struct{}

// OrderItem implements the EcommerceItemRpcImpl interface.
func (s *EcommerceItemRpcImpl) OrderItem(ctx context.Context, req *item_rpc.OrderItemRequest) (resp *item_rpc.OrderItemResponse, err error) {
	// TODO: Your code here...
	return handler.OrderItem(ctx, req)
}

// GetItem implements the EcommerceItemRpcImpl interface.
func (s *EcommerceItemRpcImpl) GetItem(ctx context.Context, req *item_rpc.GetItemRequest) (resp *item_rpc.GetItemResponse, err error) {
	// TODO: Your code here...
	return handler.GetItem(ctx, req)
}

// CreateItem implements the EcommerceItemRpcImpl
func (s *EcommerceItemRpcImpl) CreateItem(ctx context.Context, req *item_rpc.CreateItemGroupRequest) (resp *item_rpc.CreateItemGroupResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateItem implements the EcommerceItemRpcImpl interface.
func (s *EcommerceItemRpcImpl) UpdateItem(ctx context.Context, req *item_rpc.UpdateItemRequest) (resp *item_rpc.UpdateItemResponse, err error) {
	// TODO: Your code here...
	return handler.UpdateItem(ctx, req)
}

// CreateItemGroup implements the EcommerceItemRpcImpl interface.
func (s *EcommerceItemRpcImpl) CreateItemGroup(ctx context.Context, req *item_rpc.CreateItemGroupRequest) (resp *item_rpc.CreateItemGroupResponse, err error) {
	// TODO: Your code here...
	return handler.CreateItemGroup(ctx, req)
}

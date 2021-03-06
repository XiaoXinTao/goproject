// Code generated by Kitex v0.2.0. DO NOT EDIT.

package ecommerceitemrpc

import (
	"context"
	"github.com/XiaoXinTao/goproject/item_rpc/kitex_gen/xintao/project/item_rpc"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	OrderItem(ctx context.Context, req *item_rpc.OrderItemRequest, callOptions ...callopt.Option) (r *item_rpc.OrderItemResponse, err error)
	GetItem(ctx context.Context, req *item_rpc.GetItemRequest, callOptions ...callopt.Option) (r *item_rpc.GetItemResponse, err error)
	CreateItemGroup(ctx context.Context, req *item_rpc.CreateItemGroupRequest, callOptions ...callopt.Option) (r *item_rpc.CreateItemGroupResponse, err error)
	UpdateItem(ctx context.Context, req *item_rpc.UpdateItemRequest, callOptions ...callopt.Option) (r *item_rpc.UpdateItemResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kEcommerceItemRpcClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEcommerceItemRpcClient struct {
	*kClient
}

func (p *kEcommerceItemRpcClient) OrderItem(ctx context.Context, req *item_rpc.OrderItemRequest, callOptions ...callopt.Option) (r *item_rpc.OrderItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OrderItem(ctx, req)
}

func (p *kEcommerceItemRpcClient) GetItem(ctx context.Context, req *item_rpc.GetItemRequest, callOptions ...callopt.Option) (r *item_rpc.GetItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetItem(ctx, req)
}

func (p *kEcommerceItemRpcClient) CreateItemGroup(ctx context.Context, req *item_rpc.CreateItemGroupRequest, callOptions ...callopt.Option) (r *item_rpc.CreateItemGroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateItemGroup(ctx, req)
}

func (p *kEcommerceItemRpcClient) UpdateItem(ctx context.Context, req *item_rpc.UpdateItemRequest, callOptions ...callopt.Option) (r *item_rpc.UpdateItemResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateItem(ctx, req)
}

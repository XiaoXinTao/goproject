// Code generated by Kitex v0.2.0. DO NOT EDIT.
package ecommerceitemrpc

import (
	"github.com/XiaoXinTao/goproject/item_rpc/kitex_gen/xintao/project/item_rpc"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler item_rpc.EcommerceItemRpc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

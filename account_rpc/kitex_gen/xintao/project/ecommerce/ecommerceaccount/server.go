// Code generated by Kitex v0.2.0. DO NOT EDIT.
package ecommerceaccount

import (
	"account_rpc/kitex_gen/xintao/project/ecommerce"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler ecommerce.EcommerceAccount, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

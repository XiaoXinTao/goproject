package handler

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"

	"account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

func Login(ctx context.Context, req *account_rpc.LoginRequest) (resp *account_rpc.LoginResponse, err error) {
	resp = account_rpc.NewLoginResponse()
	if req.PassportUid <= 0 || len(req.PassportCode) <= 0 {
		klog.CtxErrorf(ctx, "Login handler, 参数缺少")
		return resp, errors.New("Login handler, 参数缺少")
	}

	return
}
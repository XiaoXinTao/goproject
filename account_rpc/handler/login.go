package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"account_rpc/consts"
	"account_rpc/dal"
	"account_rpc/db"
	"account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

func Login(ctx context.Context, req *account_rpc.LoginRequest) (resp *account_rpc.LoginResponse, err error) {
	resp = account_rpc.NewLoginResponse()
	if req.PassportUid <= 0 || len(req.PassportCode) <= 0 {
		klog.CtxErrorf(ctx, "Login handler, 参数缺少 ")
		return resp, errors.New("Login handler, 参数缺少 ")
	}
	if _, err = db.CheckUserIsExist(ctx, req.PassportUid, req.PassportCode); err != nil {
		klog.CtxErrorf(ctx, "Login handler, 验证账号密码错误, err = %+v", err)
		resp.BaseResp.StatusCode = consts.LoginPassportNotMatchCode
		resp.BaseResp.StatusMessage = consts.LoginPassportNotMatchMsg
		return resp, errors.New("Login handler, 验证账号密码错误 ")
	}
	accessToken := dal.GenAccessToken(ctx)
	err = dal.SaveUserToken(ctx, strconv.FormatInt(req.PassportUid, 10), strconv.FormatInt(accessToken, 10))
	if err != nil {
		klog.CtxErrorf(ctx, "Login handler, 设置AccessToken错误, err = %+v")
		resp.BaseResp.StatusCode = consts.AccessTokenSetErrCode
		resp.BaseResp.StatusMessage = consts.AccessTokenSetErrMsg
		return resp, errors.New("Login handler, 设置AccessToken错误 ")
	}
	resp.AccessToken = strconv.FormatInt(accessToken, 10)
	return resp, nil
}
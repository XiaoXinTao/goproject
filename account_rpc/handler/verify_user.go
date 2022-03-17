package handler

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"project.ecommerce.account_rpc/consts"
	"project.ecommerce.account_rpc/dal"
	"project.ecommerce.account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

func VerifyUser(ctx context.Context, req *account_rpc.VerifyUserRequest) (resp *account_rpc.VerifyUserResponse, err error) {
	resp = account_rpc.NewVerifyUserResponse()
	if req.PassportUid <= 0 || req.AccessToken <= 0 {
		klog.CtxErrorf(ctx, "VerifyUser handler 缺少参数")
		return resp, errors.New("VerifyUser hander 缺少参数")
	}
	PassportUidStr := strconv.FormatInt(req.PassportUid, 10)
	AccessTokenStr := strconv.FormatInt(req.AccessToken, 10)
	AccessTokenRds, err := dal.GetUserToken(ctx, PassportUidStr)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyUser handler, 从redis获取AccessToken错误, err = %+v", err)
		resp.BaseResp.StatusCode = consts.AccessTokenNotFoundErrCode
		resp.BaseResp.StatusMessage = consts.AccessTokenNotFoundErrMsg
		return resp, errors.New("从redis获取AccessToken错误")
	}
	if AccessTokenStr != AccessTokenRds {
		klog.CtxWarnf(ctx, "VerifyUser handler, 用户AccessToken不匹配redis里面的")
		resp.BaseResp.StatusCode = consts.AccessTokenNotMatchCode
		resp.BaseResp.StatusMessage = consts.AccessTokenNotMatchMsg
		return resp, errors.New("用户AccessToken不匹配redis里面的")
	}
	return resp, nil
}

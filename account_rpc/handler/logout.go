package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/XiaoXinTao/goproject/account_rpc/dal"
	"github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

func Logout(ctx context.Context, req *account_rpc.LogoutRequest) (resp *account_rpc.LogoutResponse, err error) {
	resp = account_rpc.NewLogoutResponse()
	if req.PassportUid <= 0 {
		klog.CtxErrorf(ctx, "Logout handler, 请求参数缺失 ")
		return resp, errors.New("Logout handler, 请求参数缺失 ")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout handler, 请求参数反序列化错误, err = %+v", err)
		return resp, errors.New("请求参数反序列化错误 系统错误")
	}
	klog.CtxInfof(ctx, "Logout handler, 请求参数反序列化 = %+v", string(JsonReq))
	if err := dal.DelUserToken(ctx, strconv.FormatInt(req.PassportUid, 10)); err != nil {
		klog.CtxErrorf(ctx, "Logout handler, 删除用户accesstoken错误, err = %+v", err)
		return resp, errors.New("删除用户accesstoken错误")
	}
	JsonResp, err := json.Marshal(resp)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout handler, 回包参数反序列化错误, err = %+v", err)
		return resp, errors.New("回包参数反序列化错误")
	}
	klog.CtxInfof(ctx, "Logout handler, 回包参数反序列化 = %+v", string(JsonResp))
	return resp, nil
}
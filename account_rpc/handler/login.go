package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/XiaoXinTao/goproject/account_rpc/consts"
	"github.com/XiaoXinTao/goproject/account_rpc/dal"
	"github.com/XiaoXinTao/goproject/account_rpc/db"
	"github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

func Login(ctx context.Context, req *account_rpc.LoginRequest) (resp *account_rpc.LoginResponse, err error) {
	resp = account_rpc.NewLoginResponse()
	if req.PassportUid <= 0 || len(req.PassportCode) <= 0 {
		klog.CtxErrorf(ctx, "Login handler, 参数缺少 ")
		return resp, errors.New("Login handler, 参数缺少 ")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "Login handler, request to json error, err = +%v", err)
		return resp, errors.New("系统内部错误")
	}
	klog.CtxInfof(ctx, "Login handler, JsonReq = %+v", string(JsonReq))
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
	JsonResp, err := json.Marshal(resp)
	if err != nil {
		klog.CtxErrorf(ctx, "Login handler, response to json error, err = +%v", err)
		return resp, errors.New("系统内部错误")
	}
	klog.CtxInfof(ctx, "Login handler, JsonReq = %+v", string(JsonResp))
	return resp, nil
}
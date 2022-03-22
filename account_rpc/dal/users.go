package dal

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
)

func SaveUserToken(ctx context.Context, PassportUid, Token string) error {
	if len(PassportUid) <= 0 || len(Token) <= 0 {
		klog.CtxErrorf(ctx, "SaveUserToken 缺少参数")
		return errors.New("SaveUserToken 缺少参数")
	}
	Key := fmt.Sprintf(keyUserToken, PassportUid)
	err := rdb.Set(PassportUid, Key, ExpirationToken).Err()
	if err != nil {
		klog.CtxErrorf(ctx, "SaveUserToken redis set error, err = %+v", err)
		return err
	}
	return nil
}

func GetUserToken(ctx context.Context, PassportUid string) (string, error) {
	if len(PassportUid) <= 0 {
		klog.CtxErrorf(ctx, "GetUserToken 缺少参数")
		return "", errors.New("GetUserToken 缺少参数")
	}
	Key := fmt.Sprintln(keyUserToken, PassportUid)
	val, err := rdb.Get(Key).Result()
	if err != nil {
		klog.CtxErrorf(ctx,"GetUserToken redis get error, err = %+v", err)
		return "", err
	}
	return val, nil
}

func DelUserToken(ctx context.Context, PassportUid string) error {
	if len(PassportUid) <= 0 {
		klog.CtxErrorf(ctx, "DelUserToken 缺少参数")
		return errors.New("DelUserToken 缺少参数")
	}
	Key := fmt.Sprintln(keyUserToken, PassportUid)
	err := rdb.Del(Key).Err()
	if err != nil {
		klog.CtxErrorf(ctx, "DelUserToken redis del error, err = %+v", err)
		return err
	}
	return nil
}

func GenAccessToken(ctx context.Context) int64 {
	return snowFlake.NextVal(ctx)
}
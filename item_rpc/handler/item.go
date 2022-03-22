package handler

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/XiaoXinTao/goproject/item_rpc/db"
	"github.com/XiaoXinTao/goproject/item_rpc/kitex_gen/xintao/project/item_rpc"
	"github.com/XiaoXinTao/goproject/item_rpc/model"
)

func OrderItem(ctx context.Context, req *item_rpc.OrderItemRequest) (resp *item_rpc.OrderItemResponse, err error) {
	resp = item_rpc.NewOrderItemResponse()
	if req.ItemId <= 0 || req.ItemCount <= 0 || req.ItemGroupId <= 0 || req.PassportUid <= 0 {
		klog.CtxErrorf(ctx, "OrderItem handler, 请求参数缺失")
		return resp, errors.New("请求参数缺失")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "OrderItem handler, 请求参数反序列化失败, err = %+v", err)
		return resp, errors.New("请求参数反序列化失败 系统错误")
	}
	klog.CtxInfof(ctx, "OrderItem handler, 请求参数反序列化 = %+v", string(JsonReq))
	Record := &model.DelieverItemsRecord{
		ItemGroupid:      req.ItemGroupId,
		ItemId:           req.ItemId,
		PassportUid:      req.PassportUid,
		DeliverItemCount: int(req.ItemCount),
		CreateTime:       time.Now(),
	}
	err = db.DeliverItem(ctx, Record)
	if err != nil {
		klog.CtxInfof(ctx, "OrderItem handler, 发货数据库异常 err = %+v", err)
		return resp, errors.New("发货数据库异常 系统错误")
	}
	JsonResp, err := json.Marshal(resp)
	klog.CtxInfof(ctx, "OrderItem handler, 回包参数反序列化失败 = %+v", string(JsonResp))
	return resp, nil
}

func GetItem(ctx context.Context, req *item_rpc.GetItemRequest) (resp *item_rpc.GetItemResponse, err error) {
	resp = item_rpc.NewGetItemResponse()
	if req.ItemGroupId <= 0 {
		klog.CtxErrorf(ctx, "GetItem handler, 请求参数缺失")
		return resp, errors.New("请求参数缺失")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "GetItem handler, 请求参数反序列化失败, err = %+v", err)
		return resp, errors.New("请求参数反序列化 系统错误")
	}
	klog.CtxInfof(ctx, "GetItem handler, 请求参数反序列化 = %+v", string(JsonReq))
	ItemList, err := db.GetItemList(ctx, req.ItemGroupId)
	if err != nil {
		klog.CtxErrorf(ctx, "GetItem handler, 获取商品数据库失败, err = %+v", err)
		return resp, errors.New("获取去商品数据库失败 系统错误")
	}
	if ItemList == nil {
		klog.CtxErrorf(ctx, "GetItem handler, 获取商品数据库失败 没找到相关记录")
		return resp, errors.New("没找到相关记录")
	}

	for _, Item := range *ItemList {
		resp.ItemList = append(resp.ItemList, &item_rpc.ItemInfo{
			ItemId:    Item.ItemId,
			ItemName:  Item.ItemName,
			ItemCount: int32(Item.ItemCount),
			IsLimited: int32(Item.IsLimited),
		})
	}
	JsonResp, err := json.Marshal(resp)
	klog.CtxInfof(ctx, "GetItem handler, 回包参数反序列化 = %+v", string(JsonResp))
	return resp, nil
}


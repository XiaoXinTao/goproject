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

// CreateItemGroup
func CreateItemGroup(ctx context.Context, req *item_rpc.CreateItemGroupRequest) (resp *item_rpc.CreateItemGroupResponse, err error) {
	resp = item_rpc.NewCreateItemGroupResponse()
	if req.ItemGroupId <= 0 || req.PassportUid <= 0 {
		klog.CtxErrorf(ctx, "CreateItemGroup handler, 请求参数缺失")
		return resp, errors.New("请求参数缺失")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateItemGroup handler, 请求参数反序列化失败 err = %+v", err)
		return resp, errors.New("请求参数反序列化失败 系统错误")
	}
	klog.CtxInfof(ctx, "CreateItemGroup handler, 请求参数反序列化 = %+v", string(JsonReq))
	ItemGroup := &model.ItemGroups{
		ItemGroupid: req.ItemGroupId,
		PassportUid: req.PassportUid,
	}
	err = db.CreateItemGroup(ctx, ItemGroup)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateItemGroup handler, 数据库创建失败 err = %+v", err)
		return resp, errors.New("数据库创建失败 系统错误")
	}
	JsonResp, err := json.Marshal(resp)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateItemGroup handler, 回包参数反序列化失败 err = %+v", err)
		return resp, errors.New("回包参数反序列化失败 系统错误")
	}
	klog.CtxInfof(ctx, "CreateItemGroup handler, 回包参数反序列化 = %+v", string(JsonResp))
	return resp, nil
}

// 只有owner商家能更新商品库存什么的
func UpdateItem(ctx context.Context, req *item_rpc.UpdateItemRequest) (resp *item_rpc.UpdateItemResponse, err error) {
	resp = item_rpc.NewUpdateItemResponse()
	if len(req.ItemList) <= 0 || req.ItemGroupId <= 0 || req.PassportUid <= 0 {
		klog.CtxErrorf(ctx, "UpdateItem handler, 请求参数缺失")
		return resp, errors.New("请求参数缺失")
	}
	JsonReq, err := json.Marshal(req)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateItem handler, 请求参数反序列化失败, err = %+v", err)
		return resp, errors.New("请求参数缺失")
	}
	klog.CtxInfof(ctx, "UpdateItem handler, 请求参数反序列化 = %+v", JsonReq)
	if IsExist, err := db.CheckItemGroupExist(ctx, req.ItemGroupId, req.PassportUid); err != nil || !IsExist {
		if err != nil {
			klog.CtxErrorf(ctx, "UpdateItem handler, 查看是否存在商家商店记录数据库错误, err = %+v", err)
			return resp, errors.New("查看是否存在商家记录数据库错误 系统错误")
		}
		if !IsExist {
			klog.CtxWarnf(ctx, "UpdateItem handler, 商家商店记录数据库无此记录")
			return resp, errors.New("请求参数为非法操作")
		}
	}
	for _, Item := range req.ItemList {
		IsExist, err := db.CheckItem(ctx, req.ItemGroupId, Item.ItemId)
		if err != nil {
			klog.CtxErrorf(ctx, "UpdateItem handler, 查询物品是否存在错误, err = %+v", err)
			return resp, errors.New("查询物品是否存在错误 系统错误")
		}
		item := &model.Items{
			ItemId:      Item.ItemId,
			ItemGroupid: req.ItemGroupId,
			ItemCount:   int(Item.ItemCount),
			ItemName:    Item.ItemName,
			IsLimited:   int(Item.IsLimited),
			PassportUid: req.PassportUid,
		}
		if IsExist {
			err = db.UpdateItemForOwner(ctx, item)
		} else {
			item.CreateTime = time.Now()
			err = db.CreateItem(ctx, item)
		}
		if err != nil {
			klog.CtxErrorf(ctx, "UpdateItem handler, 更新物品错误, err = %+v", err)
			return resp, errors.New("更新物品错误 系统错误")
		}
	}
	return resp, nil
}

package db

import (
	"context"

	"github.com/XiaoXinTao/goproject/item_rpc/model"
)

func CreateItemGroup(ctx context.Context, ItemGroup *model.ItemGroups) error {
	return getDb(ctx).Debug().Model(&model.ItemGroups{}).Create(&ItemGroup).Error
}

func CheckItemGroupExist(ctx context.Context, ItemGroupId, PassportUid int64) (bool, error) {
	var count int64 = 0
	err := getDb(ctx).Debug().Model(&model.ItemGroups{}).Where("item_groupid = ?", ItemGroupId).
		Where("passport_uid = ?", PassportUid).Count(&count).Error
	return count > 0, err
}
package db

import (
	"context"

	"gorm.io/gorm"

	"github.com/XiaoXinTao/goproject/item_rpc/model"
)

func CreateItem(ctx context.Context, Item *model.Items) error {
	return getDb(ctx).Debug().Model(&model.Items{}).Create(&Item).Error
}

func GetItemList(ctx context.Context, ItemGroupId int64) (ItemList *[]model.Items, err error) {
	err = getDb(ctx).Debug().Model(&model.Items{}).Where("item_groupid = ?", ItemGroupId).Find(&ItemList).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return ItemList, err
}

func CheckItem(ctx context.Context, ItemGroupId, ItemId int64) (bool ,error) {
	var count int64 = 0
	err := getDb(ctx).Debug().Model(&model.Items{}).Where("item_groupid = ?", ItemGroupId).
		Where("item_id = ?", ItemId).Count(&count).Error
	return count > 0, err
}

func UpdateItemForOwner(ctx context.Context, Item *model.Items) error {
	err := getDb(ctx).Debug().Model(&model.Items{}).Where("item_groupid = ?", Item.ItemGroupid).
		Where("item_id = ?", Item.ItemId).Updates(*Item).Error
	return err
}
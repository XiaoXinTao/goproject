package db

import (
	"context"

	"gorm.io/gorm"

	"github.com/XiaoXinTao/goproject/item_rpc/model"
)

func DeliverItem(ctx context.Context, Record *model.DelieverItemsRecord) error {
	return getDb(ctx).Debug().Transaction(func(tx *gorm.DB) error {
		err := getDb(ctx).Debug().Model(&model.Items{}).Where("item_groupid = ?", Record.ItemGroupid).
			Where("item_id = ?", Record.ItemId).Where("item_count >= ?", Record.DeliverItemCount).
			Update("item_count", gorm.Expr("item_count - ", Record.DeliverItemCount)).Error
		if err != nil {
			return err
		}
		err = getDb(ctx).Debug().Model(&model.DelieverItemsRecord{}).Create(&Record).Error
		if err != nil {
			return err
		}
		return nil
	})
}


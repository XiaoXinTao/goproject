// Code generated by COMMENTS_BUILD_TOOLS 2.0.68. DO NOT EDIT.
package model

import (
	"time"
)

func (Items) TableName() string {
	return "items"
}

func NewItems() *Items {
	return &Items{}
}

type Items struct {
	ItemCount   int       `gorm:"column:item_count" json:"item_count"`           // 商品剩余量
	IsLimited   int       `gorm:"column:is_limited" json:"is_limited"`           // 商品是否限量
	UniqueId    int64     `gorm:"column:unique_id;primary_key" json:"unique_id"` // id
	ItemGroupid int64     `gorm:"column:item_groupid" json:"item_groupid"`       // 商品的group
	ItemId      int64     `gorm:"column:item_id" json:"item_id"`                 // 商品的id
	PassportUid int64     `gorm:"column:passport_uid" json:"passport_uid"`       // 用户账号
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`         // create_time
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time"`         // update_time
	ItemName    string    `gorm:"column:item_name" json:"item_name"`             // 商品名称
}
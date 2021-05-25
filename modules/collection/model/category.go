package model

import (
	"gorm.io/gorm"
	"kingford-backend/global"
	"kingford-backend/utils"
)

type CollectionItem struct {
	global.BaseModel
	CollectionId string `json:"collection_id" gorm:"not null;comment:'收藏id'"`
	Name string `json:"name" gorm:"size:20;not null;comment:'名称'"`
	Description string `json:"description" gorm:"size:200;comment:'描述'"`
	Cover string `json:"cover" gorm:"comment:'封面图片'"`
	Link string `json:"link" gorm:"comment:'跳转链接地址url'"`
}


func (m *CollectionItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = utils.GenUUID()
	return
}
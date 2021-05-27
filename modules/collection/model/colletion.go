package model

import (
	"gorm.io/gorm"
	"kingford-backend/global"
	"kingford-backend/utils"
)

type Collection struct {
	CollectionCategoryId string `json:"categoryId" gorm:"not null;comment:'所属分类'"`
	Name string `json:"name" gorm:"size:20;not null;comment:'名称'"`
	Description string `json:"description" gorm:"size:200;comment:'描述'"`
	Cover string `json:"cover" gorm:"comment:'封面图片'"`
	Link string `json:"link" gorm:"comment:'跳转链接地址url'"`
	global.BaseModel
}

func (m *Collection) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = utils.GenUUID()
	return
}
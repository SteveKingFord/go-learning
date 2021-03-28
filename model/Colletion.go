package model

import (
	"gorm.io/gorm"
	"kingford-backend/utils"
)

type Collection struct {
	BaseModel
	Title string `json:"title" gorm:"size:50;not null;unique;comment:'收藏标题'"`
	CollectionItems []CollectionItem `json:"collection_items"`
}

func (m *Collection) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = utils.GenUUID()
	return
}
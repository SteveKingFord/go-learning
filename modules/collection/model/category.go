package model

import (
	"gorm.io/gorm"
	"kingford-backend/global"
	"kingford-backend/utils"
)

type CollectionCategory struct {
	Name string      `json:"name" gorm:"size:50;not null;unique;comment:'收藏标题'"`
	Collections []*Collection `json:"collections"`
	global.BaseModel
}


func (m *CollectionCategory) BeforeCreate(tx *gorm.DB) (err error) {
	m.Id = utils.GenUUID()
	return
}
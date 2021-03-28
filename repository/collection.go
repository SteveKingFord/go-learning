package repository

import (
	"gorm.io/gorm"
	"kingford-backend/global"
	"kingford-backend/model"
)

type CollectionRepository struct {
	DB      *gorm.DB
}

type CollectionRInterface interface {
	create() (*model.Collection, error)
}

func (r *CollectionRepository) Create() (*model.Collection, error) {
	entity := &model.Collection{Title: "GOLANG"}
	err := global.DB.Create(&entity).Error
	return entity, err
}

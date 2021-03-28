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

func (r *CollectionRepository) Create(collection *model.Collection) (*model.Collection, error) {
	err := global.DB.Create(collection).Error
	return collection, err
}

package repository

import (
	"gorm.io/gorm"
	"kingford-backend/model"
	"kingford-backend/utils"
)

type CollectionInterface interface {
	Create(collection *model.Collection) (*model.Collection, error)
}

type CollectionRepository struct {
	DB *gorm.DB
}

func (r *CollectionRepository) Get(id string) (*model.Collection, error) {
	var collection model.Collection
	err := r.DB.Where("id = ?", id).First(&collection).Error
	return &collection, err
}

func (r *CollectionRepository) GetList(pageIndex int, pageSize int) (collections []*model.Collection, err error) {
	limit, offset := utils.Page(pageIndex, pageSize)
	err = r.DB.Limit(limit).Offset(offset).Find(&collections).Error
	return collections, err
}

func (r *CollectionRepository) Create(collection *model.Collection) (*model.Collection, error) {
	err := r.DB.Create(collection).Error
	return collection, err
}

func (r *CollectionRepository) Update(id string, collection *model.Collection) (*model.Collection, error) {
	err := r.DB.Model(&model.Collection{}).Where("id = ?", id).Updates(collection).Error
	return collection, err
}

func (r *CollectionRepository) Delete(id string) error{
	var collection model.Collection
	err := r.DB.Where("id = ?", id).Delete(&collection).Error
	return err
}
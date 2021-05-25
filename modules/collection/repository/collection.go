package repository

import (
	"gorm.io/gorm"
	model2 "kingford-backend/modules/collection/model"
	"kingford-backend/utils"
)

type CollectionInterface interface {
	GetList(pageIndex int, pageSize int) (collections []*model2.Collection, err error)
	Get(id string) (*model2.Collection, error)
	Create(collection *model2.Collection) (*model2.Collection, error)
	Update(id string, collection *model2.Collection) (*model2.Collection, error)
	Delete(id string) error
}

type CollectionRepository struct {
	DB *gorm.DB
}


func (r *CollectionRepository) GetList(pageIndex int, pageSize int) (collections []*model2.Collection, err error) {
	limit, offset := utils.Page(pageIndex, pageSize)
	err = r.DB.Preload("CollectionItems").Order("title").Limit(limit).Offset(offset).Find(&collections).Error
	return collections, err
}


func (r *CollectionRepository) Get(id string) (*model2.Collection, error) {
	var collection model2.Collection
	err := r.DB.Where("id = ?", id).First(&collection).Error
	return &collection, err
}

func (r *CollectionRepository) Create(collection *model2.Collection) (*model2.Collection, error) {
	err := r.DB.Omit("CollectionItems").Create(collection).Error
	return collection, err
}

func (r *CollectionRepository) Update(id string, collection *model2.Collection) (*model2.Collection, error) {
	err := r.DB.Model(&model2.Collection{}).Where("id = ?", id).Updates(collection).Error
	return collection, err
}

func (r *CollectionRepository) Delete(id string) error{
	var collection model2.Collection
	err := r.DB.Where("id = ?", id).Delete(&collection).Error
	return err
}
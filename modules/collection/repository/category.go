package repository

import (
	"gorm.io/gorm"
	model2 "kingford-backend/modules/collection/model"
	"kingford-backend/utils"
)

type CollectionItemRepository struct {
	DB *gorm.DB
}

func (r *CollectionItemRepository) GetList(pageIndex int, pageSize int) (CollectionItems []*model2.CollectionItem, err error) {
	limit, offset := utils.Page(pageIndex, pageSize)
	err = r.DB.Order("name").Limit(limit).Offset(offset).Find(&CollectionItems).Error
	return CollectionItems, err
}

func (r *CollectionItemRepository) Get(id string) (CollectionItem *model2.CollectionItem, err error) {
	err = r.DB.Where("id = ?", id).First(&CollectionItem).Error
	return CollectionItem, err
}

func (r *CollectionItemRepository) Create(CollectionItem *model2.CollectionItem) (*model2.CollectionItem, error) {
	err := r.DB.Create(CollectionItem).Error
	return CollectionItem, err
}

func (r *CollectionItemRepository) Update(id string, CollectionItem *model2.CollectionItem) (*model2.CollectionItem, error) {
	err := r.DB.Model(&model2.CollectionItem{}).Where("id=?", id).Updates(&CollectionItem).Error
	return CollectionItem, err
}

func (r *CollectionItemRepository) Delete(id string) error {
	var CollectionItem model2.CollectionItem
	return r.DB.Where("id = ?", id).Delete(&CollectionItem).Error
}

package repository

import (
	"gorm.io/gorm"
    "kingford-backend/modules/collection/model"
	"kingford-backend/utils"
)

type CollectionCategoryRepository struct {
	DB *gorm.DB
}

func (r *CollectionCategoryRepository) GetList(pageIndex int, pageSize int) (CollectionCategories []*model.CollectionCategory, err error) {
	limit, offset := utils.Page(pageIndex, pageSize)
	err = r.DB.Preload("Collections").Order("name").Limit(limit).Offset(offset).Find(&CollectionCategories).Error
	return CollectionCategories, err
}

func (r *CollectionCategoryRepository) Get(id string) (CollectionCategory *model.CollectionCategory, err error) {
	err = r.DB.Where("id = ?", id).First(&CollectionCategory).Error
	return CollectionCategory, err
}

func (r *CollectionCategoryRepository) Create(CollectionCategory *model.CollectionCategory) (*model.CollectionCategory, error) {
	err := r.DB.Create(CollectionCategory).Error
	return CollectionCategory, err
}

func (r *CollectionCategoryRepository) Update(id string, CollectionCategory *model.CollectionCategory) (*model.CollectionCategory, error) {
	err := r.DB.Model(&model.CollectionCategory{}).Where("id=?", id).Updates(&CollectionCategory).Error
	return CollectionCategory, err
}

func (r *CollectionCategoryRepository) Delete(id string) error {
	var CollectionCategory model.CollectionCategory
	return r.DB.Where("id = ?", id).Delete(&CollectionCategory).Error
}

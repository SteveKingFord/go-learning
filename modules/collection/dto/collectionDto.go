package dto

import "kingford-backend/modules/collection/model"

type Collection struct {
	Id string `json:"id"`
	CategoryId string `json:"categoryId"`
	Name string `json:"name" `
	Description string `json:"description" `
	Cover string `json:"cover"`
	Link string `json:"link"`
}


func BuildCollectionDto(item *model.Collection) *Collection {
	return &Collection{
		Id:item.Id,
		CategoryId: item.CollectionCategoryId,
		Name: item.Name,
		Description: item.Description,
		Cover:item.Cover,
		Link:item.Link,
	}
}


func BuildCollectionsDto(items []*model.Collection) (list []*Collection){
	for _, item := range items {
		elem := BuildCollectionDto(item)
		list = append(list,elem)
	}
	return list
}
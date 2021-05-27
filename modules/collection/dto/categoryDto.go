package dto

import "kingford-backend/modules/collection/model"

type CategoryDto struct {
	Id string                `json:"id"`
	Name string              `json:"name"`
	Collections []*Collection `json:"collections"`
}




func BuildCategoryDto(item *model.CollectionCategory) *CategoryDto{

	return  &CategoryDto{
		Id:item.Id,
		Name:item.Name,
		Collections: BuildCollectionsDto(item.Collections),
	}
}


func BuildCategoriesDto(items []*model.CollectionCategory)(list []*CategoryDto){
	for _, item := range items {
		elem := BuildCategoryDto(item)
		list = append(list,elem)
	}
	return list
}
package dto

import "kingford-backend/model"

// Menu 数据
type Menu struct {
	Id    string `json:"id"`
	Title string `json:"name"`
}

// BuildCollection 单个序列化
func BuildCollection(item model.Collection) Menu {
	return Menu{
		Id: item.Id,
		Title: item.Title,
	}
}

// BuildBuildCollections 序列化列表
func BuildBuildCollections(list []model.Collection) (items []Menu) {
	for _, item := range list {
		value := BuildCollection(item)
		items = append(items, value)
	}
	return items
}

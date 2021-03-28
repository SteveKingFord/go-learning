package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/model"
	"kingford-backend/repository"
)

type UpdateItemService struct {
	CollectionId string `json:"collectionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Cover        string `json:"cover"`
	Link         string `json:"link"`
}

func (s *UpdateItemService) Update(id string) *dto.Response {
	repo := repository.CollectionItemRepository{DB: global.DB}

	var entity = &model.CollectionItem{
		CollectionId: s.CollectionId,
		Name:         s.Name,
		Description:  s.Description,
		Cover:        s.Cover,
		Link:         s.Link,
	}
	item, err := repo.Update(id, entity)

	if err != nil {
		return &dto.Response{
			Status: 400,
			Data:   item,
			Msg:    err.Error(),
		}
	}
	return &dto.Response{
		Status: 200,
		Data:   item,
		Msg:    "success",
	}
}

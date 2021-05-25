package category

import (
	"kingford-backend/global"
	model2 "kingford-backend/modules/collection/model"
	repository2 "kingford-backend/modules/collection/repository"
)

type UpdateItemService struct {
	CollectionId string `json:"collectionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Cover        string `json:"cover"`
	Link         string `json:"link"`
}

func (s *UpdateItemService) Update(id string) *global.Response {
	repo := repository2.CollectionItemRepository{DB: global.DB}

	var entity = &model2.CollectionItem{
		CollectionId: s.CollectionId,
		Name:         s.Name,
		Description:  s.Description,
		Cover:        s.Cover,
		Link:         s.Link,
	}
	item, err := repo.Update(id, entity)

	if err != nil {
		return &global.Response{
			Status: 400,
			Data:   item,
			Msg:    err.Error(),
		}
	}
	return &global.Response{
		Status: 200,
		Data:   item,
		Msg:    "success",
	}
}

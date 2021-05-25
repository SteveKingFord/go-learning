package collection

import (
	"kingford-backend/global"
	model2 "kingford-backend/modules/collection/model"
	repository2 "kingford-backend/modules/collection/repository"
)

type UpdateService struct {
	Title string `json:"title" form:"title"`
}

func (s *UpdateService) Update(id string) *global.Response {
	repo := repository2.CollectionRepository{DB: global.DB}

	entity := &model2.Collection{
		Title: s.Title,
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

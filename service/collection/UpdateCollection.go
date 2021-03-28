package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/model"
	"kingford-backend/repository"
)

type UpdateService struct {
	Title string `json:"title" form:"title"`
}

func (s *UpdateService) Update(id string) *dto.Response {
	repo := repository.CollectionRepository{DB: global.DB}

	entity := &model.Collection{
		Title: s.Title,
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

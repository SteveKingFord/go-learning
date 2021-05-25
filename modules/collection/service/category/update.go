package category

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/model"
	"kingford-backend/modules/collection/repository"
)

type UpdateService struct {
	Name        string `json:"name"`
}

func (s *UpdateService) Update(id string) *global.Response {
	repo := repository.CollectionCategoryRepository{DB: global.DB}

	var entity = &model.CollectionCategory{
		Name: s.Name,
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

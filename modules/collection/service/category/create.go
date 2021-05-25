package category

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/model"
	"kingford-backend/modules/collection/repository"
	"net/http"
)

type CreateService struct {
	Name        string `json:"name"`
}

func (s *CreateService) Create() *global.Response {
	resp := repository.CollectionCategoryRepository{
		DB: global.DB,
	}

	entity := &model.CollectionCategory{
		Name: s.Name,
	}

	item, err := resp.Create(entity)

	if err != nil {
		return &global.Response{
			Status: http.StatusFailedDependency,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}

	return &global.Response{
		Status: http.StatusOK,
		Data:   item,
	}
}

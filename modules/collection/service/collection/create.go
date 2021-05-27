package collection

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/model"
	"kingford-backend/modules/collection/repository"
	"net/http"
)

type CreateService struct {
	CategoryId string `json:"categoryId"`
	Name string `json:"name" form:"name"`
	Cover string `json:"cover"`
	Link string `json:"link"`
	Description string `json:"description"`
}

func (s *CreateService) Create() *global.Response {

	resp := repository.CollectionRepository{
		DB: global.DB,
	}

	var entity = &model.Collection{
		CollectionCategoryId: s.CategoryId,
		Name:     s.Name,
		Cover:s.Cover,
		Link:s.Link,
		Description: s.Description,
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

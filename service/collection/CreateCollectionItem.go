package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/model"
	"kingford-backend/repository"
	"net/http"
)

type CreateItemService struct {
	CollectionId string `json:"collectionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Cover        string `json:"cover"`
	Link         string `json:"link"`
}

func (s *CreateItemService) Create() *dto.Response {
	resp := repository.CollectionItemRepository{
		DB: global.DB,
	}

	var entity = &model.CollectionItem{
		CollectionId: s.CollectionId,
		Name:         s.Name,
		Description:  s.Description,
		Cover:        s.Cover,
		Link:         s.Link,
	}

	item, err := resp.Create(entity)

	if err != nil {
		return &dto.Response{
			Status: http.StatusFailedDependency,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}

	return &dto.Response{
		Status: http.StatusOK,
		Data:   item,
	}
}

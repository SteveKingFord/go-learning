package category

import (
	"kingford-backend/global"
	model2 "kingford-backend/modules/collection/model"
	repository2 "kingford-backend/modules/collection/repository"
	"net/http"
)

type CreateItemService struct {
	CollectionId string `json:"collectionId"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Cover        string `json:"cover"`
	Link         string `json:"link"`
}

func (s *CreateItemService) Create() *global.Response {
	resp := repository2.CollectionItemRepository{
		DB: global.DB,
	}

	var entity = &model2.CollectionItem{
		CollectionId: s.CollectionId,
		Name:         s.Name,
		Description:  s.Description,
		Cover:        s.Cover,
		Link:         s.Link,
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

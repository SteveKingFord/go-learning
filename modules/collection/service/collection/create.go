package collection

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/model"
	"kingford-backend/modules/collection/repository"
	"net/http"
)

type CreateService struct {
	Name string `json:"name" form:"name"`
}

func (s *CreateService) Create() *global.Response {

	resp := repository.CollectionRepository{
		DB: global.DB,
	}

	//uuid := utils.GenUUID()

	var entity = &model.Collection{
		//BaseModel: model.BaseModel{Id:  utils.GenUUID()},
		Name:     s.Name,
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

package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/model"
	"kingford-backend/repository"
	"net/http"
)

type CreateService struct {
	Title string `json:"title" form:"title"`
}

func (s *CreateService) Create() *dto.Response {

	resp := repository.CollectionRepository{
		DB: global.DB,
	}

	//uuid := utils.GenUUID()

	var entity = &model.Collection{
		//BaseModel: model.BaseModel{Id:  utils.GenUUID()},
		Title:     s.Title,
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

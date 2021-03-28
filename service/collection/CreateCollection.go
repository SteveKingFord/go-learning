package collection

import (
	"kingford-backend/dto"
	"kingford-backend/model"
	"kingford-backend/repository"
	"kingford-backend/utils"
	"net/http"
)

type CreateService struct {
	Title string `json:"title" form:"title"`
}

func (s *CreateService) Create() *dto.Response {
	resp := repository.CollectionRepository{}
	uuid, _ := utils.GenUUID()

	var entity = &model.Collection{
		BaseModel: model.BaseModel{Id: uuid},
		Title:     s.Title,
	}

	items, err := resp.Create(entity)

	if err != nil {
		return &dto.Response{
			Status: http.StatusFailedDependency,
			Msg:    "创建失败",
			Error:  err.Error(),
		}
	}

	return &dto.Response{
		Status: http.StatusOK,
		Data:   items,
	}
}

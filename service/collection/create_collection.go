package collection

import (
	"kingford-backend/dto"
	"kingford-backend/repository"
	"net/http"
)

type CreateService struct {
	Title string `json:"title" form:"title"`
}

func(s *CreateService) Create() *dto.Response  {
	resp := repository.CollectionRepository{}
	items,err :=resp.Create()

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
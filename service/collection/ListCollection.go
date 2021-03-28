package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/repository"
)

type ListService struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

func (s *ListService) GetList() *dto.Response {
	repo := repository.CollectionRepository{DB: global.DB}
	item, err := repo.GetList(s.PageIndex, s.PageSize)

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

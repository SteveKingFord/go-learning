package collection

import (
	"kingford-backend/global"
	 "kingford-backend/modules/collection/repository"
)

type ListService struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

func (s *ListService) GetList() *global.Response {
	repo := repository.CollectionRepository{DB: global.DB}
	item, err := repo.GetList(s.PageIndex, s.PageSize)

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

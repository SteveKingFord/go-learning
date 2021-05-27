package category

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/dto"
	"kingford-backend/modules/collection/repository"
	"net/http"
)

type ListService struct {
	PageIndex int `json:"pageIndex" form:"pageIndex"`
	PageSize  int `json:"pageSize" form:"pageSize"`
}

func (s *ListService) GetList() *global.Response {
	repo := repository.CollectionCategoryRepository{DB: global.DB}
	entity, err := repo.GetList(s.PageIndex, s.PageSize)
	list := dto.BuildCategoriesDto(entity)

	if err != nil {
		return &global.Response{
			Status: http.StatusInternalServerError,
			Data:  list,
			Msg:    err.Error(),
		}
	}

	return &global.Response{
		Status: http.StatusOK,
		Data:   list,
		Msg:    "success",
	}
}

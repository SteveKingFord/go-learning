package category

import (
	"kingford-backend/global"
	repository2 "kingford-backend/modules/collection/repository"
)

type GetItemService struct {
}

func (s *GetItemService) Get(id string) *global.Response {
	repo := repository2.CollectionItemRepository{DB: global.DB}
	item, err := repo.Get(id)

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

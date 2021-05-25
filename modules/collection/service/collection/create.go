package collection

import (
	"kingford-backend/global"
	model2 "kingford-backend/modules/collection/model"
	repository2 "kingford-backend/modules/collection/repository"
	"net/http"
)

type CreateService struct {
	Title string `json:"title" form:"title"`
}

func (s *CreateService) Create() *global.Response {

	resp := repository2.CollectionRepository{
		DB: global.DB,
	}

	//uuid := utils.GenUUID()

	var entity = &model2.Collection{
		//BaseModel: model.BaseModel{Id:  utils.GenUUID()},
		Title:     s.Title,
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

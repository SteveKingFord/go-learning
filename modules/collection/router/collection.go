package router

import (
	"github.com/gin-gonic/gin"
	"kingford-backend/modules/collection/api"
)

func RegisterCollectionRouter(router *gin.RouterGroup) {
	r := router.Group("")
	{
		r.GET("/collection", api.GetList)
		r.GET("/collection/:id", api.Get)
		r.POST("/collection", api.Create)
		r.PUT("/collection/:id", api.Update)
		r.DELETE("/collection/:id", api.Delete)

		r.GET("/collection-category", api.GetItemList)
		r.GET("/collection-category/:id", api.GetItem)
		r.POST("/collection-category", api.CreateItem)
		r.PUT("/collection-category/:id", api.UpdateItem)
		r.DELETE("/collection-category/:id", api.DeleteItem)
	}
}

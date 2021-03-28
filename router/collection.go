package router

import (
	"github.com/gin-gonic/gin"
	api "kingford-backend/api/v1/collection"
)

func RegisterCollectionRouter(router *gin.RouterGroup) {
	r := router.Group("")
	{
		r.GET("/collection", api.GetList)
		r.GET("/collection/:id", api.Get)
		r.POST("/collection", api.Create)
		r.PUT("/collection/:id", api.Update)
		r.DELETE("/collection/:id", api.Delete)

		r.GET("/collection-item", api.GetItemList)
		r.POST("/collection-item", api.CreateItem)
		r.PUT("/collection-item/:id", api.UpdateItem)
	}
}

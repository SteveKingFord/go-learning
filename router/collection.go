package router

import (
	"github.com/gin-gonic/gin"
	api "kingford-backend/api/v1/collection"
)

func RegisterCollectionSysRouter(router *gin.RouterGroup)  {
	r := router.Group("collection")
	{
		//r.GET("/", api.GetList)
		//r.GET("/:id", sys_user.Get)
		r.POST("/", api.Create)
		//r.PUT("/:id", sys_user.Put)
		//r.DELETE("/:id", sys_user.Delete)
	}
}
package router

import (
	"kingford-backend/api/v1/sys_user"
	"github.com/gin-gonic/gin"
)



func RegisterSysUserRouter(router *gin.RouterGroup)  {
	r := router.Group("system")
	{
		r.GET("/user", sys_user.GetList)
		r.GET("/user/:id", sys_user.Get)
		r.PUT("/user/:id", sys_user.Put)
		r.DELETE("/user/:id", sys_user.Delete)
	}
}
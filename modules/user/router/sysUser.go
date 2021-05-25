package router

import (
	"github.com/gin-gonic/gin"
	"kingford-backend/modules/user/api/sys_user"
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
/*
 * @Author: kingford
 * @Date: 2021-03-24 16:58:35
 * @LastEditTime: 2021-03-24 17:09:33
 */
package router

import (
	"kingford-backend/api/v1/sys_user"
	"github.com/gin-gonic/gin"
)

// RegisterSysBaseRouter 注册系统基础路由
func RegisterSysBaseRouter(router *gin.RouterGroup)  {
	r := router.Group("system")
	{
		r.POST("/register", sys_user.Register)
		r.POST("/login", sys_user.Login)
	}
}
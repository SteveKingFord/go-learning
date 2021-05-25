/*
 * @Author: kingford
 * @Date: 2021-03-24 16:58:35
 * @LastEditTime: 2021-03-24 17:09:33
 */
package router

import (
	"github.com/gin-gonic/gin"
	"kingford-backend/modules/user/api/sys_user"
)

// RegisterSysBaseRouter 注册系统基础路由
func RegisterSysBaseRouter(router *gin.RouterGroup)  {
	r := router.Group("system")
	{
		r.POST("/register", sys_user.Register)
		r.POST("/login", sys_user.Login)
	}
}
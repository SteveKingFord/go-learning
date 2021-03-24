package initialize

import (
	_ "kingford-backend/docs"
	"kingford-backend/middleware"
	"kingford-backend/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterRouter()  *gin.Engine{
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	// 配置跨越
	r.Use(middleware.Cors())

	// 配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 不需要jwt验证
	publicGroup := r.Group("")
	{
		router.RegisterSysBaseRouter(publicGroup)
	}

	// 需要jwt授权
	privateGroup := r.Group("")
	{
		router.RegisterSysUserRouter(privateGroup)
	}

	return r
}

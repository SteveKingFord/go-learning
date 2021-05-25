package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "kingford-backend/docs"
	"kingford-backend/middleware"
	collection "kingford-backend/modules/collection/router"
	user "kingford-backend/modules/user/router"
)

func RegisterRouter()  *gin.Engine{
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	// 配置跨越
	r.Use(middleware.Cors())

	// 配置swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 不需要jwt验证
	publicGroup := r.Group("/v1/api")
	{
		user.RegisterSysBaseRouter(publicGroup)
		collection.RegisterCollectionRouter(publicGroup)
	}

	// 需要jwt授权
	privateGroup := r.Group("/v1/api")
	{
		user.RegisterSysUserRouter(privateGroup)
	}

	return r
}

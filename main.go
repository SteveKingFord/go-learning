package main

import (
	"kingford-backend/core"
	"kingford-backend/global"
	"kingford-backend/initialize"
)


// @title Swagger kingford backend API
// @version 1.0
// @description This is a kingford backend  server Petstore server.

// @contact.name API Support
// @contact.email stevekingford@gmail.com

// @securityDefinitions.apikey Bearer
// @type apiKey
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	// 初始化Viper配置文件
	global.Viper = core.Viper()
	// 初始化zap日志库
	global.Log = core.Zap()
	// 初始化gorm连接数据库
	global.DB = initialize.Gorm()
	if global.DB != nil {
		initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	// 注册路由
	core.RunWindowsServer()
}

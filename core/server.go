package core

import (
	"fmt"
	"kingford-backend/global"
	"kingford-backend/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	router := initialize.RegisterRouter()
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 Gin-Template
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
	`, address)

	global.Log.Error(s.ListenAndServe().Error())
}

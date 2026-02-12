package main

import (
	"mygo_bangforai/pkg/router"
	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/logger"
	"mygo_bangforai/pkg/utils"
)

func main() {
	err:=config.InitConfig()
	if err!=nil{
		panic(err)
	}

	err=logger.InitLogger()
	if err!=nil{
		panic(err)
	}
	logger.Info("日志初始化完成")

	err=utils.InitJWT()
	if err!=nil{
		logger.Fatalf("初始化JWT失败: %v", err)
	}
	logger.Info("JWT初始化完成")
	
	err=config.InitMySQL()
	if err!=nil{
		logger.Fatalf("初始化MySQL失败: %v", err)
	}
	logger.Info("MySQL初始化完成")
	
	r,err := router.SetupRouter()
	if err!=nil{
		logger.Fatalf("初始化路由失败: %v", err)
	}
	logger.Info("路由初始化完成")
	serverPort := ":" + config.GetServerPort() + ""// 从配置中获取端口号
	r.Run(serverPort)
	logger.Info("服务器启动完成，监听端口: " + serverPort)
}

//http://localhost:8080
//git push origin main

package main

import (
	"mygo_bangforai/pkg/router"
	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/logger"
	"mygo_bangforai/pkg/utils"
)

func main() {
	logger.InitLogger()// 初始化日志
	logger.Info("程序启动")
	
	config.InitConfig()// 初始化配置
	logger.Info("配置初始化完成")
	
	utils.InitJWT()// 初始化JWT
	logger.Info("JWT初始化完成")
	
	config.InitMySQL()// 初始化 MySQL 数据库
	logger.Info("数据库初始化完成")
	
	r := router.SetupRouter()
	serverPort := ":" + config.GetServerPort() + ""// 从配置中获取端口号
	logger.Info("服务器启动", logger.Zap.String("port", serverPort))
	r.Run(serverPort)
}

//http://localhost:8080
//git push origin main

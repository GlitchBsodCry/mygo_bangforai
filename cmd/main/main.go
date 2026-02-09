package main

import (
	"mygo_bangforai/pkg/router"
	"mygo_bangforai/pkg/config"
	"mygo_bangforai/pkg/logger"
	"mygo_bangforai/pkg/utils"
)

func main() {
	logger.InitLogger()// 初始化日志
	
	config.InitConfig()// 初始化配置
	
	utils.InitJWT()// 初始化JWT
	
	config.InitMySQL()// 初始化 MySQL 数据库
	
	r := router.SetupRouter()
	serverPort := ":" + config.GetServerPort() + ""// 从配置中获取端口号
	r.Run(serverPort)
}

//http://localhost:8080
//git push origin main

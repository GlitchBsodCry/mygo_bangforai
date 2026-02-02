package main

import (
	"mygo_bangforai/pkg/router"
	"mygo_bangforai/pkg/config"
)

func main() {
	config.InitConfig()// 初始化配置
	r := router.SetupRouter()
	serverPort := ":" + config.GetServerPort() + ""// 从配置中获取端口号
	r.Run(serverPort)
}

//http://localhost:8080

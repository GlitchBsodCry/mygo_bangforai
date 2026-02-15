package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	//"mygo_bangforai/api/errors"
	"mygo_bangforai/internal/control"
	"mygo_bangforai/pkg/middleware"
	"mygo_bangforai/pkg/interfacer"
)
var logger = interfacer.GetLogger()
func SetupRouter() (*gin.Engine, error) {
	r := gin.Default()
	r.Use(middleware.Recovery())
	logger.Info("创建路由成功", zap.String("path", "/"))
	ruser := r.Group("/user")
	{
		ruser.POST("/register", control.Register)
		ruser.POST("/login", control.Login)
	}
	return r,nil
}

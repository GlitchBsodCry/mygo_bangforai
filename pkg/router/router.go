package router

import (
	"github.com/gin-gonic/gin"
	
	"mygo_bangforai/internal/control"
	"mygo_bangforai/pkg/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Recovery())
	ruser := r.Group("/user")
	{
		ruser.POST("/register", control.Register)
		ruser.POST("/login", control.Login)
	}
	return r
}

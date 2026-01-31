package control

import (
	"mygo_bangforai/api/model"
	"mygo_bangforai/api/error/response"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c,  err.Error())
		return
	}
	response.Success(c, gin.H{
		"username": req.Username,
		"email":    req.Email,
	})
}

func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}
	response.Success(c, gin.H{
		"username": req.Username,
	})
}
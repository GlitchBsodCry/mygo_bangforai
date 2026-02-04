package control

import (
	"mygo_bangforai/api/error/response"
	"mygo_bangforai/api/model"
	"mygo_bangforai/internal/service"
	"mygo_bangforai/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c,  err.Error())
		return
	}

	user, err := service.Register(req)
	if err != nil {
		response.Error(c, response.InternalError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ParamError(c, err.Error())
		return
	}

	user, err := service.Login(req)
	if err != nil {
		response.Error(c, response.Unauthorized, err.Error())
		return
	}

	// 生成JWT token
    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        response.Error(c, response.InternalError, "Failed to generate token")
        return
    }

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"token":    token,
	})
}
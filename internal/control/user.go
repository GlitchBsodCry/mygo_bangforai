package control

import (
	"mygo_bangforai/api/errors"
	"mygo_bangforai/api/model"
	"mygo_bangforai/internal/service"
	//"mygo_bangforai/pkg/logger"
	"mygo_bangforai/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.ParamError(c,  err.Error())
		return
	}

	user, err := service.Register(req)
	if err != nil {
		errors.Error(c, errors.InternalError, err.Error())
		return
	}

	errors.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.ParamError(c, err.Error())
		return
	}

	user, err := service.Login(req)
	if err != nil {
		errors.Error(c, errors.Unauthorized, err.Error())	
		return
	}

	// 生成JWT token
    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        errors.Error(c, errors.InternalError, "Failed to generate token")
        return
    }

	errors.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"token":    token,
	})
}
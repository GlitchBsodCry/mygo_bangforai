package middleware

import (
	"mygo_bangforai/api/errors"
	"mygo_bangforai/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			errors.Error(c, errors.Unauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		tokenString := authHeader
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			errors.Error(c, errors.Unauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
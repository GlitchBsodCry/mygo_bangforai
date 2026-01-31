package middleware

import (
	//"net/http"

	"mygo_bangforai/api/error/response"
	
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Error(c, response.InternalError, "internal error")
				c.Abort()
			}
		}()
		c.Next() 
	}
}
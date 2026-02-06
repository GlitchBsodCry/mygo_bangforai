package middleware

import (
	//"net/http"

	"mygo_bangforai/api/errors"
	
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errors.Error(c, errors.InternalError, "internal error")
				c.Abort()
			}
		}()
		c.Next() 
	}
}
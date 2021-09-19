package middlewares

import (
	"github.com/gin-gonic/gin"
)

func JsonHeaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()
	}
}

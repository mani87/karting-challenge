package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthentic() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")
		if key != "apitest" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

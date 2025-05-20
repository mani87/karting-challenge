package middleware

import (
	"github.com/gin-gonic/gin"
	"karting-challenge/internal/config"
	"net/http"
)

func IsAuthentic() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("x-api-key")
		if key != config.ApiKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

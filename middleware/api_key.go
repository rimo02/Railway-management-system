package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// You need to protect all the admin API endpoints with an API key that will be known only to you and the admin so that no one can add false data to your system.
func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != "only_admin_knows_this_api_key" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized access"})
			return
		}
		c.Next()
	}
}

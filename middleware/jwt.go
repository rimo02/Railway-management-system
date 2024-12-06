package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Railway-management-system/utils"
)

// For booking a seat and getting specific booking details, you need to send the Authorization Token received in the login endpoint.
func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            return
        }
        claims, err := utils.ValidateJWT(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }
        c.Set("userID", claims["userID"])
        c.Set("role", claims["role"])
        c.Next()
    }
}

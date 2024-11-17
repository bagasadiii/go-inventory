package middleware

import (
	"inventory-app/internal/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer "){
			c.JSON(http.StatusUnauthorized, gin.H{"error":"authorization required"})
			c.Abort()
			return
		}
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"authorization required"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"authorization required"})
			c.Abort()
			return
		}
		username,err := helper.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"token invalid"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}

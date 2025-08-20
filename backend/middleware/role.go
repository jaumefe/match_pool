package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "role not found in context"})
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to register matches"})
			return
		}

		c.Next()
	}
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists || userRole != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireAdmin(c *gin.Context) {
	userRole, exists := c.Get("userRole")
	if !exists || userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized – Admins only"})
		c.Abort()
		return
	}
	c.Next()
}

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc{
	return func(c *gin.Context){
		role, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "role not found",
			})
			return 
		}

		userRole, ok := role.(string)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid role",
			})
			return
		}

		for _, allowed := range allowedRoles {
			if userRole == allowed {
				c.Next()
				return 
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error" : "insufficient permission",
		})
	}
}
package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := ""

		if authHeader != "" {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if tokenString == "" {
			cookieToken, err := c.Cookie("access_token")
			if err == nil {
				tokenString = cookieToken
			}
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)

		sub, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		userId := uint(sub)

		c.Set("userID", userId)
		c.Set("role", claims["role"])

		c.Next()

	}
}

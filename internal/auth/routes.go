package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func RegisterRoutes(api *gin.RouterGroup, handler *Handler) {

	auth := api.Group("/auth")

	auth.POST("/login", handler.Login)

	protected := auth.Group("/")
	protected.Use(AuthMiddleware())

	protected.POST("/logout", handler.Logout)
	protected.GET("/me", handler.Me)
	protected.POST("/create-user", middleware.RequireRoles("god"), handler.CreateUser)
}

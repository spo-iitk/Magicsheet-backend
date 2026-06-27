package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(api *gin.RouterGroup){
	
	auth := api.Group("/auth")
	
	auth.POST("/login",Login)
	auth.POST("/signup", SignUp)
	auth.POST("/logout", Logout)
}

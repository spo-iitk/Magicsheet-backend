package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type singUp struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Login Endpoint reached",
	})
}

func SignUp(c *gin.Context){
	var req singUp
	
	err := c.BindJSON(&req)
if err != nil {
    fmt.Println("BindJSON error:", err)
    return
}

fmt.Println("Works!", req)
	c.JSON(http.StatusOK, gin.H{
		"message" : "signup endpoint reached",
	})
}

func Logout(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "logout endpoint reached",
	})
}
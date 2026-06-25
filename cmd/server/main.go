package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		godotenv.Load(".env") // fallback if run from project root
	}

	r := gin.Default()

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		panic(err)
	}

	r.GET("/health",func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status" : "ok",
		})
	})
	
	r.Run(":8080")
}




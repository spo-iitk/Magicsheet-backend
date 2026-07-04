package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		godotenv.Load(".env")
	}

	//creating router
	r := gin.Default()

	//cors 
	r.Use(cors.New(middleware.CORS()))
	api := r.Group("/api")

	//db init
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	//depedency injection
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)
	auth.RegisterRoutes(api, handler)

	//migration
	if err := database.AutoMigrate(db); err != nil {
		panic(err)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}

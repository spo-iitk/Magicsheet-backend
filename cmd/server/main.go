package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
	"github.com/spo-iitk/Magicsheet-backend/internal/rc"
	"github.com/spo-iitk/Magicsheet-backend/internal/sync"
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

	//database init
	pibsDB, err := database.InitPIBSDB()
	if err != nil {
		panic(err)
	}

	rasRCDB, err := database.InitRASRCDB()
	if err != nil {
		panic(err)
	}

	rasApplicationDB, err := database.InitRASApplicationDB()
	if err != nil {
		panic(err)
	}
	rasStudentDB, err := database.InitRASStudentDB()
	if err != nil {
		panic(err)
	}

	//dependecy injection
	//auth module
	repo := auth.NewRepository(pibsDB)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)
	auth.RegisterRoutes(api, handler)

	//sync module
	syncRepo := sync.NewRepository(pibsDB)
	rasRepo := sync.NewRASrepository(rasRCDB, rasApplicationDB, rasStudentDB)
	syncService := sync.NewService(syncRepo, rasRepo)
	syncHandler := sync.NewHandler(syncService)
	sync.RegisterRoutes(api, syncHandler)

	// recruitment cycles
	rcRepo := rc.NewRepository(pibsDB)
	rcService := rc.NewService(rcRepo)
	rcHandler := rc.NewHandler(rcService)
	rc.RegisterRoutes(api, rcHandler)

	if err := database.AutoMigrate(pibsDB); err != nil {
		panic(err)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}

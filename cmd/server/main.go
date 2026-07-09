package main

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spo-iitk/Magicsheet-backend/internal/auth"
	"github.com/spo-iitk/Magicsheet-backend/internal/database"
	"github.com/spo-iitk/Magicsheet-backend/internal/magicsheet"
	"github.com/spo-iitk/Magicsheet-backend/internal/middleware"
	"github.com/spo-iitk/Magicsheet-backend/internal/rc"
	"github.com/spo-iitk/Magicsheet-backend/internal/sync"
)

func main() {
	if err := loadEnvFile(); err != nil {
		log.Printf("warning: could not load .env file: %v", err)
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

	magicRepo := magicsheet.NewRepository(pibsDB)
	magicService := magicsheet.NewService(magicRepo)
	magicHandler := magicsheet.NewHandler(magicService)

	magicsheet.RegisterRoutes(api, magicHandler, magicRepo)

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

func loadEnvFile() error {
	_, currentFile, _, ok := runtime.Caller(0)
	if ok {
		projectRootEnv := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(currentFile))), ".env")
		if err := godotenv.Load(projectRootEnv); err == nil {
			return nil
		}
	}

	if err := godotenv.Load(".env"); err == nil {
		return nil
	}

	return godotenv.Load("../../.env")
}

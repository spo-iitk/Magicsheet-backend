package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    fmt.Println("Connected!")
	return DB, nil
}
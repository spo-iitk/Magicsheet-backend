package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPIBSDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
}

func InitRASDB(dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("RAS_POSTGRES_HOST"),
		os.Getenv("RAS_POSTGRES_USER"),
		os.Getenv("RAS_POSTGRES_PASSWORD"),
		dbName,
		os.Getenv("RAS_POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	fmt.Printf("\n!!--------Connnected to Ras %s--------!!\n", dbName)

	return db, nil

}

func InitRASRCDB() (*gorm.DB, error) {
	return InitRASDB("rc")
}

func InitRASApplicationDB() (*gorm.DB, error) {
	return InitRASDB("application")
}
func InitRASStudentDB() (*gorm.DB, error) {
	return InitRASDB("student")
}

package config

import (
	"fmt"
	"mini-project/entity"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func SetupConnectionDatabase() *gorm.DB {
	// If use localhost uncomment this section and change sslmode from require to disable

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("Failed to create a connection to your database")
	}

	db = conn
	db.AutoMigrate(&entity.Applicant{}, &entity.Employee{}, &entity.Jobexperience{})
	return db
}

func CloseConnectionDatabase(db *gorm.DB) {
	dbPostgre := db.DB()
	
	err := dbPostgre.Close()
	if err != nil {
		panic("Failed to close connection to your database")
	}
}
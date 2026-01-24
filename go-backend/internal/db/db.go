package db

import (
	"log"
	"fmt"
	"os"
    "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {  
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")	
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",dbUser,dbPass,dbHost,dbPort,dbName,
	)
	//dsn := "root:password@tcp(localhost:3306)/godb?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = database
	log.Println("Database connected successfully")
}

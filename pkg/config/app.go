package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db * gorm.DB
)

func Connect(){
	err := godotenv.Load()
    if err != nil {
        log.Panic("Error loading .env file")
    }

    // Retrieve the values from environment variables
    host := os.Getenv("MYSQL_HOST")
    user := os.Getenv("MYSQL_USER")
    password := os.Getenv("MYSQL_PASSWORD")
    

    // Create the DSN (Data Source Name)
    dsn := fmt.Sprintf("%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host)

    // Open the connection using Gorm and MySQL driver
    d, err := gorm.Open("mysql", dsn)
    if err != nil {
        log.Panic("Failed to connect to database:", err)
    }

    db = d
    fmt.Println("Database connection successful")
}

func GetDB() *gorm.DB{
	return db
}
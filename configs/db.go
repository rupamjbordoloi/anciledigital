package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db != nil {
		println("returning old connection")
		return db
	}
	err := godotenv.Load()
	if err != nil {
		panic("Failed to connect to the database")
	}
	println("creating new connection")
	username := os.Getenv("DB_USERNAME")
	pwd := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=true", username, pwd, database)
	// Connect to the MySQL database using the MySQL dialect
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}

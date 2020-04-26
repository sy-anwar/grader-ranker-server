package config

import (
	"os"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName   := os.Getenv("db_name")
	dbHost   := os.Getenv("db_host")
	dbPort   := os.Getenv("db_port")
	
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	database, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("db connected")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
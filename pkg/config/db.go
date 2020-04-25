package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() {
	dbUri := "root:@tcp(127.0.0.1:3306)/test_db?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dbUri)
	defer database.Close()
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
package config

import (
	"fmt"

	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/?parseTime=True")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS orders_by")
	db.Exec("USE orders_by")
	db.AutoMigrate(&models.Item{}, &models.Order{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

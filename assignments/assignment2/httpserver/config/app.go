package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// func initDB() {
// 	var err error
// 	db, err = gorm.Open("mysql", "root:@tcp(localhost:8080)/?parseTime=True")

// 	if err != nil {
// 		fmt.Println(err)
// 		panic("failed to connect database")
// 	}
// 	db.Exec("USE orders_by")
// 	db.AutoMigrate(&models.Order{}, &models.Item{})
// }

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:8080)/orders_by")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

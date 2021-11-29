package driver

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	uri := "root:@tcp(127.0.0.1:3306)/antonio?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMigrate()
}

func initMigrate() {
	models := registerEntities()
	for _, model := range models {
		DB.AutoMigrate(&model.Model)
	}
}
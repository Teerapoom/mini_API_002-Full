package database

import (
	"github.com/teerapoom/miniapi_version2/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/moive_miniapi002?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.Movie{}, &model.Director{})
	Db = db
}

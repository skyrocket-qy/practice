package orm

import (
	"sql-practice/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("./example.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		panic("Failed to perform auto migration")
	}
}

func GetDB() *gorm.DB {
	return db
}

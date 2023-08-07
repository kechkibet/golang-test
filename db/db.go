package db

import (
	"cms-tizi/db/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Run() {
	db, err := gorm.Open(sqlite.Open("my-database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
	err = db.AutoMigrate(&schema.Todo{})
	if err != nil {
		return
	}
}

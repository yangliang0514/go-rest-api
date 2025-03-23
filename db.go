package main

import (
	"github.com/yangliang0514/go-rest-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Event{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

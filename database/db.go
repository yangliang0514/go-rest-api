package database

import (
	"github.com/yangliang0514/go-rest-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database/app.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Event{})

	return GetDB(db)
}

func GetDB(db ...*gorm.DB) *gorm.DB {
	if len(db) > 0 && db[0] != nil {
		DB = db[0]
	}
	return DB
}

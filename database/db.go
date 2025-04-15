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

	SetDB(db)
	SetupMigrations()

	return DB
}

func SetDB(db *gorm.DB) {
	DB = db
}

func SetupMigrations() {
	DB.AutoMigrate(&models.Event{})
	DB.AutoMigrate(&models.User{})
}

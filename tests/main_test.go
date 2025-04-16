package tests

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/database"
	"github.com/yangliang0514/go-rest-api/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var server *gin.Engine

func TestMain(m *testing.M) {
	server = setupServer()
	setupTestDB()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupServer() *gin.Engine {
	return router.RegisterRoutes(gin.Default())
}

func setupTestDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("failed to connect database")
	}

	database.SetDB(db)
	database.SetupMigrations()
}

package database

import (
	"os"
	"task-manager/packages/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	home, _ := os.UserHomeDir()
	
	os.MkdirAll(home+"/.local/share/task-manager", os.ModePerm)

	db, err := gorm.Open(sqlite.Open(home+"/.local/share/task-manager/data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Task{})
}

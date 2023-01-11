package config

import (
	"github.com/mattbaxter689/projectStore/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Implement a function that checks if there's a database, and creates if
// missing

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("error connecting to the database")
	}
	
	db.AutoMigrate(&models.Message{}, &models.Project{})

	return db
}

package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Implement a function that checks if there's a database, and creates if
// missing

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("projectStore.db"), &gorm.Config{})
	if err != nil {
		panic("error connecting to the database")
	}

	return db
}

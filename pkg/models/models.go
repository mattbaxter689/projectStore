package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Messages []Message `gorm:"foreignKey:ProjectID"`
}

type Message struct {
	gorm.Model
	Comment string
	ProjectID uint 
}

package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Messages []Message
}

type Message struct {
	gorm.Model
	Comment string
	ProjectID uint 

}

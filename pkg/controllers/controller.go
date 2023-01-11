package controllers

import (
	"fmt"

	// "github.com/mattbaxter689/projectStore/pkg/models"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
	Messages []Message `gorm:"foreignKey:ProjectID"`
}

type Message struct {
	// gorm.Model
	Comment string
	ProjectID uint
}

type GormRepository struct {
	DB *gorm.DB
}

//HasProject Used to determine if a new project should be created
func (g *GormRepository) HasProject() bool {
	if projects, _ := g.GetAllProjects(); len(projects) == 0 {
		return false
	}
	return true
}

func (g *GormRepository) CreateNewProject() error {
	msg := Message{Comment: "testing the db and what this does"}
	proj := Project{Name: "DATA"}
	proj.Messages = []Message{msg}
	res := g.DB.Create(&proj)

	return  res.Error
}

// GetAllProjects gets all active projects currently found in the db
func (g *GormRepository) GetAllProjects() ([]Project, error) {
	var projects []Project

	if err := g.DB.Preload("Messages").Find(&projects).Error; err != nil {
		return projects, fmt.Errorf("Table is empty: %v", err)
	}

	return projects, nil
	
}

//The entry handlers
func (g *GormRepository) AddEntryByProject() error {
	msg := Message{Comment: "adding a new comment to test how this works"}
	var project Project
	if err := g.DB.Where("id = ?", 1).First(&project).Error; err != nil {
		panic("Problem extracting project")
	}

	project.Messages = append(project.Messages, msg)

	err := g.DB.Save(&project).Error
	if err != nil {
		panic("Problem saving data to database")
	}
	
	return err
}

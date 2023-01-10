package controllers

import (
	"fmt"

	"github.com/mattbaxter689/projectStore/pkg/models"
	"gorm.io/gorm"
)

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

// GetAllProjects gets all active projects currently found in the db
func (g *GormRepository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project

	if err := g.DB.Preload("Messages").Find(&projects).Error; err != nil {
		return projects, fmt.Errorf("Table is empty: %v", err)
	}

	return projects, nil
	
}

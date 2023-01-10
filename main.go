package main

import (
	// "flag"
	"fmt"

	"github.com/mattbaxter689/projectStore/config"
	"github.com/mattbaxter689/projectStore/pkg/controllers"
)

func main() {
	// by default a help flag is given to us so that is not needed
	db := config.Connect()	
	// fmt.Println("db connected: %s", db) 
	pr := controllers.GormRepository{DB: db}

	projects := pr.HasProject()
	if !projects {
		fmt.Println("You do not currently have any projects")
	}
}

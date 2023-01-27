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

	pr := controllers.GormRepository{DB: db}

	projects := pr.HasProject()
	if !projects {
		err := pr.CreateNewProject()
		if err != nil {
			panic("Error entering value into database")
		}
	}

	if err := pr.AddEntryByProject(); err != nil {
		panic("Problem adding entry to the database")
	}

	proj, err := pr.GetAllProjects()
	if err != nil {
		panic("Error retrieiving data")
	}
	fmt.Println(proj)

}

package main

import (
	"fmt"
	// "flag"
	"github.com/mattbaxter689/projectStore/config"
)

func main() {
	// by default a help flag is given to us so that is not needed
	config.Connect()	
	fmt.Println("db connected")
}

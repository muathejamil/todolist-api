package main

import (
	"fmt"
	"todolist-api/config"
)

// main project entry point.
func main() {
	configuration, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("error with getting configuration")
		return
	}
	InitServer(configuration)
}

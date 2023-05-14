package main

import "todolist-api/storage"

// init inits the connection to db.
func init() {
	//resources.LoadEnvVariables()
	storage.ConnectToDb()
}

// main project entry point.
func main() {
	InitServer()
}

//How to change line 20, 21 https://github.com/golang/go/issues/29261

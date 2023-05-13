package main

import "todolist-api/storage"

func init() {
	//resources.LoadEnvVariables()
	storage.ConnectToDb()
}

func main() {
	InitServer()
}

//How to change line 20, 21 https://github.com/golang/go/issues/29261

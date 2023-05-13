package migrate

import (
	"todolist-api/models"
	"todolist-api/storage"
)

func init() {
	//resources.LoadEnvVariables()
	storage.ConnectToDb()
}

func main() {
	storage.DB.AutoMigrate(&models.Todo{})
}

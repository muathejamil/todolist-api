package migrate

import (
	"todolist-api/models"
	"todolist-api/resources"
)

func init() {
	resources.LoadEnvVariables()
	resources.ConnectToDb()
}

func main() {
	resources.DB.AutoMigrate(&models.Todo{})
}

package migrate

import (
	"todolist-api/models"
	"todolist-api/storage"
)

// init connects to db.
func init() {
	//resources.LoadEnvVariables()
	storage.ConnectToDb()
}

// main migration main driver.
func main() {
	storage.DB.AutoMigrate(&models.Todo{})
}

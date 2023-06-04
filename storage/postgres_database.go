package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
	"todolist-api/config"
)

var onlyOnce sync.Once

// GetDBConnection connects to db.
func GetDBConnection(database config.DataBase) *gorm.DB {
	var db *gorm.DB
	onlyOnce.Do(func() {
		var err error
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", database.Host, database.UserName, database.Password, database.Name, database.Port, database.SSL)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database")
		}
		return
	})
	return db
}

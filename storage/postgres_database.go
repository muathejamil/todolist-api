package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var onlyOnce sync.Once

// Use sync once
//https://www.socketloop.com/tutorials/golang-how-to-run-your-code-only-once-with-sync-once-object

// GetDBConnection connects to db.
func GetDBConnection() *gorm.DB {
	var db *gorm.DB
	onlyOnce.Do(func() {
		var err error
		dsn := "host=drona.db.elephantsql.com user=nboreafj password=Cv6qr6J6Fml74tgt7fib3qVkp5Z-zzDC dbname=nboreafj port=5432 sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database")
		}
		return
	})
	return db
}

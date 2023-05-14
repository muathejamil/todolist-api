package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var onlyOnce sync.Once
var DB *gorm.DB

// Use sync once
//https://www.socketloop.com/tutorials/golang-how-to-run-your-code-only-once-with-sync-once-object

// ConnectToDb connects to db.
func ConnectToDb() {
	var err error
	onlyOnce.Do(func() {
		dsn := "host=drona.db.elephantsql.com user=nboreafj password=Cv6qr6J6Fml74tgt7fib3qVkp5Z-zzDC dbname=nboreafj port=5432 sslmode=disable"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("Failed to connect to database")
		}
	})

}

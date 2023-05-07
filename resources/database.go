package resources

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// Use sync once
//https://www.socketloop.com/tutorials/golang-how-to-run-your-code-only-once-with-sync-once-object

func ConnectToDb() {
	dsn := "host=drona.db.elephantsql.com user=nboreafj password=Cv6qr6J6Fml74tgt7fib3qVkp5Z-zzDC dbname=nboreafj port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}

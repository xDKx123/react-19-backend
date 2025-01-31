package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func ConnectLocalDatabase() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println("Connecting to the database...")
	log.Println(dsn)

	var db *gorm.DB
	var err error

	// Retry logic: attempt to connect to the database up to 10 times
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to the database successfully!")
			break
		}
		log.Printf("Attempt %d: Failed to connect to the database. Retrying in 2 seconds...\n", i)
		time.Sleep(2 * time.Second)
	}

	if db != nil {
		log.Println("Database connection is ready to use.")
		return db, nil
	}

	return nil, err
}

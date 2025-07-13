package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("Waiting for Database connection Start up...")
	time.Sleep(5 * time.Millisecond)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening Database : %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the Database : %v", err)
	}

	fmt.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatal("Error Closing The Database: %V", err)
	}
}

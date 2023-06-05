package datastore

import (
	"database/sql"
	_ "database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbInstance *sql.DB

// Initialise and Connect to DB
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("CONNSTR")
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = db
}

func CloseDB() {
	defer dbInstance.Close()
}

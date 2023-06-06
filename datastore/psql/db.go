package datastore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

// Close DB connection
func CloseDB() {
	defer dbInstance.Close()
}

// login function
func Login(username string, password string) bool {
	rows, error := dbInstance.Query("Select password from users where username=($1)", username)
	if error != nil {
		fmt.Println(error)
	}
	defer rows.Close()
	var pass string
	for rows.Next() {
		err := rows.Scan(&pass)
		if err != nil {
			log.Println(err)
		}
	}
	if error == nil {
		err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
		if err == nil {
			return true
		} else {
			return false
		}
	}
	return false
}

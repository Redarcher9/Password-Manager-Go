package data

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

type credential struct {
	UserId   int
	CredDesc string
	CredId   string
	CredPass string
}

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
	dbInstance = db
	rows, err := dbInstance.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var userID int
	var username string
	var pass string
	for rows.Next() {
		err := rows.Scan(&userID, &username, &pass)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(userID, username, pass)
	}
}

func SignUp(username string, password string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 19)
	if err != nil {
		log.Print(err)
	}
	encryptedPass := string(bytes)
	_, userError := dbInstance.Query("Select * from users where username='($1)'", username)
	fmt.Println(userError)
	if userError != nil {
		fmt.Println(userError)
	} else {
		_, error := dbInstance.Exec("Insert into users (username,password) values ($1,$2)", username, encryptedPass)
		if error != nil {
			fmt.Println(error)
		}
	}
}

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

func GetCredentials(id int) []credential {
	rows, error := dbInstance.Query("Select * from credentials where userid=$1", id)
	if error != nil {
		fmt.Println(error)
	}
	defer rows.Close()
	var credlist []credential
	for rows.Next() {
		var c credential
		err := rows.Scan(&c.UserId, &c.CredDesc, &c.CredId, &c.CredPass)
		if err != nil {
			log.Println(err)
		}
		credlist = append(credlist, c)
	}
	return credlist
}

func GetId(username string) int {
	rows, error := dbInstance.Query("Select userid from users where username=$1", username)
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close()
	var userID int
	for rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			log.Println(err)
		}
	}
	return userID
}

func AddCredentials(userID int, credDesc string, credID string, credPass string) bool {
	_, error := dbInstance.Exec("Insert into credentials (userid,cred_desc,cred_id,cred_pass) values ($1,$2,$3,$4)", userID, credDesc, credID, credPass)
	if error != nil {
		fmt.Println(error)
		return false
	}
	return true
}

func DeleteCredential(userID int, credDesc string) bool {
	_, error := dbInstance.Exec("Delete from credentials where cred_desc=$1 and userid=$2", credDesc, userID)
	if error != nil {
		fmt.Println(error)
		return false
	}
	return true
}

func UpdateCredentials(userID int, credDesc string, credID string, credPass string) bool {
	_, error := dbInstance.Exec("Update credentials SET cred_desc = $2,cred_id = $3,cred_pass = $4 where userid=$1 and cred_desc = $2", userID, credDesc, credID, credPass)
	if error != nil {
		fmt.Println(error)
		return false
	}
	return true
}

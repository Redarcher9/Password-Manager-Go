package datastore

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Inmem struct {
}

func (I *Inmem) Login(username string, password string) bool {
	var realusername = "snehit"
	var realpassword = "password@123"
	if realusername == username {
		log.Println("username does not exist")
	}
	err := bcrypt.CompareHashAndPassword([]byte(realpassword), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}

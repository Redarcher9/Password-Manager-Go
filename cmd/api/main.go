package main

import (
	"fmt"
	"password-manager/data"
)

func main() {
	data.ConnectDB()
	fmt.Println("Welcome to password manager. This app is written in GO!")
	r := setupRouter()
	r.Run(":8080")
	data.CloseDB()
}

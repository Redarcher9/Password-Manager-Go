package handlers

import (
	"fmt"
	"net/http"
	"password-manager/data"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		//bind the JSON payload
		err := c.ShouldBindJSON(&u)
		//check for errors while binding
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//authenticate with the database
		data.SignUp(u.Username, u.Password)
		fmt.Println(u.Username)
		//Authorised Successfully, return 200
		c.JSON(http.StatusOK, gin.H{"status": "user created"})
	}
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		//bind JSON payload
		err := c.ShouldBindJSON(&u)
		//check for errors while binding
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//authenticate with the database
		loginResult := data.Login(u.Username, u.Password)
		if loginResult {
			c.JSON(http.StatusOK, gin.H{
				"status": "Login Successful",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "Login Unsuccessful, credentials incorrect",
			})
			return
		}
	}
}

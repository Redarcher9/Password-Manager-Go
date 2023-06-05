package handlers

import (
	"net/http"
	"password-manager/data"

	"github.com/gin-gonic/gin"
)

func LoginHandler2() gin.HandlerFunc {
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

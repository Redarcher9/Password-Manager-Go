package handlers

import (
	"fmt"
	"net/http"
	"password-manager/data"

	"github.com/gin-gonic/gin"
)

type CredStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	CredDesc string `json:"creddesc"`
	CredId   string `json:"credid"`
	CredPass string `json:"credpass"`
}

func GetCredsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u User
		//bind JSON payload
		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not Bind JSON",
			})
		}
		loginResult := data.Login(u.Username, u.Password)
		if loginResult {
			credlist := data.GetCredentials(data.GetId(u.Username))
			c.JSON(http.StatusOK, gin.H{
				"cred": credlist,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "login failed",
			})
			return
		}
	}
}

func AddCredsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cbody CredStruct
		//bind JSON payload
		err := c.ShouldBindJSON(&cbody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not Bind JSON",
			})
		}
		loginResult := data.Login(cbody.Username, cbody.Password)
		if loginResult {
			userid := data.GetId(cbody.Username)
			credResult := data.AddCredentials(userid, cbody.CredDesc, cbody.CredId, cbody.CredPass)
			if !credResult {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Could not add the user credentials",
				})
				return
			}
			credlist := data.GetCredentials(userid)
			c.JSON(http.StatusOK, gin.H{
				"cred": credlist,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "login failed",
			})
			return
		}
	}
}

func DeleteCredsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cbody CredStruct
		//bind JSON payload
		err := c.ShouldBindJSON(&cbody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not Bind JSON",
			})
		}
		loginResult := data.Login(cbody.Username, cbody.Password)
		if loginResult {
			userid := data.GetId(cbody.Username)
			fmt.Println(userid)
			credResult := data.DeleteCredential(userid, cbody.CredDesc)
			if !credResult {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Could not delete the user credentials",
				})
				return
			}
			credlist := data.GetCredentials(userid)
			c.JSON(http.StatusOK, gin.H{
				"cred": credlist,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "login failed",
			})
			return
		}
	}
}

func UpdateCredsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cbody CredStruct
		//bind JSON payload
		err := c.ShouldBindJSON(&cbody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not Bind JSON",
			})
		}
		loginResult := data.Login(cbody.Username, cbody.Password)
		if loginResult {
			userid := data.GetId(cbody.Username)
			credResult := data.UpdateCredentials(userid, cbody.CredDesc, cbody.CredId, cbody.CredPass)
			if !credResult {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Could not delete the user credentials",
				})
				return
			}
			credlist := data.GetCredentials(userid)
			c.JSON(http.StatusOK, gin.H{
				"cred": credlist,
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "login failed",
			})
			return
		}
	}
}

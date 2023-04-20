package main

import (
	"password-manager/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	//initialise Gin
	r := gin.Default()
	//Initialise Routes and redirect to handler functions
	r.GET("/ping", handlers.PingHandler())
	r.POST("/signup", handlers.SignUpHandler())
	r.POST("/login", handlers.LoginHandler())
	r.GET("/credentials", handlers.GetCredsHandler())
	r.POST("/credentials", handlers.AddCredsHandler())
	r.DELETE("/credentials", handlers.DeleteCredsHandler())
	r.PUT("/credentials", handlers.UpdateCredsHandler())
	return r
}

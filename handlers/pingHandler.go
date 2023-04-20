package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Backend is UP",
		})
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "auth healthy"})
	})

	r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler)

	r.Run(":8000")
}

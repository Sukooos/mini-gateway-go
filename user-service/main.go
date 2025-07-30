package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "User Service Healthy",
		})
	})

	r.GET("/profile", UserProfileHandler) // Endpoint to get user profile
	r.PUT("/profile", UserUpdateHandler)   // Endpoint to update user profile

	r.Run(":8001") // Start the server on port 8001
}
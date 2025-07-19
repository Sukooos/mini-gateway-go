package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Billing Service Healthy",
		})
	})

	r.GET("/list", BillingListHandler)
	r.POST("/create", BillingCreateHandler)

	r.Run(":8002") // Start the server on port 8002
}
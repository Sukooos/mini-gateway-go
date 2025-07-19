package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	InitEnv()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Healthy",
		})
	})

	api := r.Group("/api/auth") 
	{
		api.POST("/register", ProxyRegister)
		api.POST("/login", ProxyLogin)
	}

	apiUser := r.Group("/api/user")
	apiUser.Use(JWTAuthMiddleware())
	{
		apiUser.GET("/profile", ProxyUserProfile) // Endpoint untuk mendapatkan profil user
		apiUser.PUT("/profile", ProxyUserUpdate) // Endpoint untuk update profil user
	}
	
	apiBilling := r.Group("/api/billing")
	apiBilling.Use(JWTAuthMiddleware())
	{
		apiBilling.GET("/list", ProxyBillingList)
		apiBilling.POST("/create", ProxyBillingCreate)
	}

	r.Run(":8080") // Start the server on port 8080
}
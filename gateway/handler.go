package main

import (
	"net/http"
	"os"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	
)

var authServiceURL string
var billingServiceURL string

func InitEnv() {
	authServiceURL = os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "http://localhost:8080" // Default value if not set (for local development)
	}

	billingServiceURL := os.Getenv("BILLING_SERVICE_URL")
	if billingServiceURL == "" {
		billingServiceURL = "http://localhost:8002" // Default value if not set (for local development)
	}
}

// Proxy POST /api/auth/register ke Auth Service
func ProxyRegister(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		SetBody(c.Request.Body).
		SetHeader("Content-Type", "application/json").
		Post(authServiceURL + "/register")
	
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Failed to connect to auth service",
		})
		return
	}
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// Proxy POST /api/auth/login ke Auth Service
func ProxyLogin(c *gin.Context) {
	client := resty.New()
	resp, err := client.R().
		SetBody(c.Request.Body).
		SetHeader("Content-Type", "application/json").
		Post(authServiceURL + "/login")
	
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Failed to connect to Auth Service",
		})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func ProxyBillingList(c *gin.Context) {
	client := resty.New()
	// Inject user id dan role dari context ke header
	userID := c.GetInt("user_id")
	role, _ := c.Get("role")

	resp, err := client.R().
		SetHeader("X-User-ID", toString(userID)).
		SetHeader("X-User-Role", toString(role)).
		Get(billingServiceURL + "/list")

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to billing service"})
		return
	}
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}


// Proxy POST /api/billing/create
func ProxyBillingCreate(c *gin.Context) {
	client := resty.New()

	userID := c.GetInt("user_id")
	role, _ := c.Get("role")

	// Copy body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-User-ID", toString(userID)).
		SetHeader("X-User-Role", toString(role)).
		SetBody(bodyBytes).
		Post(billingServiceURL + "/create")

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to billing service"})
		return
	}
	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

// Helper convert interface{} to string (safe)
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int:
		return fmt.Sprintf("%d", val)
	default:
		return ""
	}
}
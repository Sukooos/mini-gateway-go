package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
)

var users = []User{
	{ID: 1, Email: "raja@mail.com", Name: "Raja", Role: "user", CreatedAt: "2025-07-19T00:00:00Z"},
	{ID: 2, Email: "admin@mail.com", Name: "Admin", Role: "admin", CreatedAt: "2025-07-19T00:00:00Z"},
}

// GET /profile
func UserProfileHandler(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in header"})
		return
	}
	for _, u := range users {
		if u.ID == userID {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// PUT /profile (update name, email)
func UserUpdateHandler(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in header"})
		return
	}
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	for i, u := range users {
		if u.ID == userID {
			users[i].Name = req.Name
			users[i].Email = req.Email
			c.JSON(http.StatusOK, users[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}


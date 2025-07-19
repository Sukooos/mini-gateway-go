package main

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"strconv"
)

var billing = []Billing{}
var nextID = 1

func getUserID(c *gin.Context) (int, bool) {
	uid := c.GetHeader("X-User-ID")
	if uid == "" {
		return 0, false
	}
	id, err := strconv.Atoi(uid)
	return id, err == nil
}

// GET /list — List semua billing milik user
func BillingListHandler(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id not found in header"})
		return
	}
	var userBills []Billing
	for _, b := range billing {
		if b.UserID == userID {
			userBills = append(userBills, b)
		}
	}
	c.JSON(http.StatusOK, userBills)
}

// POST /create — Hanya admin (nanti cek role di header)
func BillingCreateHandler(c *gin.Context) {
	role := c.GetHeader("X-User-Role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can create billing"})
		return
	}

	var req struct {
		UserID      int     `json:"user_id" binding:"required"`
		Amount      float64 `json:"amount" binding:"required"`
		Description string  `json:"description" binding:"required"`
		DueDate     string  `json:"due_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	newBill := Billing{
		ID:          nextID,
		UserID:      req.UserID,
		Amount:      req.Amount,
		Description: req.Description,
		Status:      "pending",
		CreatedAt:   time.Now().Format(time.RFC3339),
		DueDate:     req.DueDate,
	}
	nextID++
	billing = append(billing, newBill)

	c.JSON(http.StatusCreated, newBill)
}

package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// Helper get user_id from header
func getUserID(c *gin.Context) (int, bool) {
	uid := c.GetHeader("X-User-ID")
	if uid == "" {
		return 0, false
	}
	id, err := strconv.Atoi(uid)
	return id, err == nil
}

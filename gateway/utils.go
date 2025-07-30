package main

import (
	"fmt"
)

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
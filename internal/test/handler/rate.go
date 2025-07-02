package handler

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Rate(c *gin.Context) {
	time.Sleep(1 * time.Second / 2)
	// Simulate some work that might fail
	if rand.Intn(1000)%100 == 0 { // Simulate 1% failure rate
		time.Sleep(1 * time.Second / 2)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "simulated failure"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

package handlers

import (
	"money_transfer_system/database"
	"money_transfer_system/queue"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	DB     *database.Database
	Rabbit *queue.RabbitMQ
}

func NewHealthHandler(db *database.Database, rabbit *queue.RabbitMQ) *HealthHandler {
	return &HealthHandler{
		DB:     db,
		Rabbit: rabbit,
	}
}

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	// Check DB Connection
	err := h.DB.Conn.Ping()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "DB unavailable", "error": err.Error()})
		return
	}

	// Check RabbitMQ Connection
	if h.Rabbit.Conn.IsClosed() {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "RabbitMQ unavailable"})
		return
	}

	// If all checks pass
	c.JSON(http.StatusOK, gin.H{"status": "Healthy"})
}

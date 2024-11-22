package handlers

import (
	"money_transfer_system/database"
	"money_transfer_system/models"
	"money_transfer_system/queue"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionHandler(database *database.Database, rabbitMQ *queue.RabbitMQ) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Transaction
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Publish transaction to RabbitMQ
		err := rabbitMQ.PublishTransaction(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue transaction"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction enqueued"})
	}
}

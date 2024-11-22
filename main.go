package main

import (
	"log"
	"money_transfer_system/config"
	"money_transfer_system/database"
	"money_transfer_system/handlers"
	"money_transfer_system/queue"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configurations
	cfg := config.LoadConfig()

	// Initialize Database
	database, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Initialize RabbitMQ
	rabbitMQ, err := queue.NewRabbitMQ(cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Start Consumer
	go queue.StartConsumer(database, rabbitMQ)

	// Start API Server
	StartServer(database, rabbitMQ)
}

func StartServer(database *database.Database, rabbitMQ *queue.RabbitMQ) {
	r := gin.Default()

	healthHandlers := handlers.NewHealthHandler(database, rabbitMQ)
	// Routes
	r.GET("/health", healthHandlers.CheckHealth)
	r.POST("/transaction", handlers.CreateTransactionHandler(database, rabbitMQ))

	// Start server
	r.Run(":8080")
}

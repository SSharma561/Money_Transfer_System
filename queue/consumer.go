package queue

import (
	"encoding/json"
	"log"
	"money_transfer_system/database"
	"money_transfer_system/models"
	"money_transfer_system/services"
)

func StartConsumer(database *database.Database, rabbitMQ *RabbitMQ) {
	msgs, err := rabbitMQ.Channel.Consume(
		rabbitMQ.Queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	service := services.NewTransactionService(database)

	for msg := range msgs {
		var transaction models.Transaction
		if err := json.Unmarshal(msg.Body, &transaction); err != nil {
			log.Printf("Failed to deserialize transaction: %v", err)
			continue
		}

		err := service.ProcessTransaction(&transaction)
		if err != nil {
			log.Printf("Failed to process transaction: %v", err)
		} else {
			log.Printf("Transaction processed successfully: %+v", transaction)
		}
	}
}

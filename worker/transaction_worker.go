package worker

import (
	"encoding/json"
	"log"
	"money_transfer_system/models"
	"money_transfer_system/services"

	"github.com/streadway/amqp"
)

type TransactionWorker struct {
	service *services.TransactionService
}

func NewTransactionWorker(service *services.TransactionService) *TransactionWorker {
	return &TransactionWorker{service: service}
}

func (w *TransactionWorker) Start(queueName string, conn *amqp.Connection) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a RabbitMQ consumer: %v", err)
	}

	for msg := range msgs {
		var req models.Transaction
		if err := json.Unmarshal(msg.Body, &req); err != nil {
			log.Printf("Invalid transaction message: %v", err)
			continue
		}

		err = w.service.ProcessTransaction(&req)
		if err != nil {
			log.Printf("Failed to process transaction: %v", err)
		} else {
			log.Printf("Transaction processed: %+v", req)
		}
	}
}

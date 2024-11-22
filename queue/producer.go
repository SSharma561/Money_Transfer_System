package queue

import (
	"encoding/json"
	"fmt"
	"money_transfer_system/models"

	"github.com/streadway/amqp"
)

func (r *RabbitMQ) PublishTransaction(transaction *models.Transaction) error {
	body, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to serialize transaction: %v", err)
	}

	return r.Channel.Publish("", r.Queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

// func (r *RabbitMQ) PublishTransaction(senderID, receiverID int, amount float64) error {
// 	body := fmt.Sprintf("%d,%d,%.2f", senderID, receiverID, amount)
// 	return r.Channel.Publish("", r.Queue.Name, false, false, amqp.Publishing{
// 		ContentType: "text/plain",
// 		Body:        []byte(body),
// 	})
// }

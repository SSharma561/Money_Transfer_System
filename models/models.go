package models

type Transaction struct {
	ID         int     `json:"id"`
	SenderID   int     `json:"sender_id"`
	ReceiverID int     `json:"receiver_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"` // e.g., "PENDING", "SUCCESS", "FAILED"
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

type Client struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type TransactionRequest struct {
	SenderID   int     `json:"sender_id" binding:"required"`
	ReceiverID int     `json:"receiver_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required,gt=0"`
}

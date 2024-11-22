package services

import (
	"errors"
	"money_transfer_system/database"
	"money_transfer_system/models"
)

type TransactionService struct {
	DB *database.Database
}

func NewTransactionService(db *database.Database) *TransactionService {
	return &TransactionService{DB: db}
}

// CreateTransaction enqueues a new transaction after validations
func (s *TransactionService) CreateTransaction(senderID, receiverID int, amount float64) error {
	// Validate sender has enough balance
	sender, err := s.DB.GetClientByID(senderID)
	if err != nil {
		return err
	}
	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	// Ensure receiver exists
	_, err = s.DB.GetClientByID(receiverID)
	if err != nil {
		return err
	}

	// Create a transaction record
	transaction := models.Transaction{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Amount:     amount,
		Status:     "PENDING",
	}
	return s.DB.CreateTransaction(&transaction)
}

// ProcessTransaction updates balances and marks transaction as completed
func (s *TransactionService) ProcessTransaction(transaction *models.Transaction) error {
	err := s.DB.UpdateBalances(transaction.SenderID, transaction.ReceiverID, transaction.Amount)
	if err != nil {
		return err
	}

	// Mark transaction as SUCCESS
	return s.DB.UpdateTransactionStatus(transaction.ID, "SUCCESS")
}

package database

import "money_transfer_system/models"

func (db *Database) GetClientByID(clientID int) (*models.Client, error) {
	row := db.Conn.QueryRow("SELECT id, name, balance, created_at, updated_at FROM clients WHERE id = $1", clientID)
	client := models.Client{}
	err := row.Scan(&client.ID, &client.Name, &client.Balance, &client.CreatedAt, &client.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (db *Database) CreateTransaction(transaction *models.Transaction) error {
	_, err := db.Conn.Exec(`
		INSERT INTO transactions (sender_id, receiver_id, amount, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`, transaction.SenderID, transaction.ReceiverID, transaction.Amount, transaction.Status)
	return err
}

func (db *Database) UpdateTransactionStatus(transactionID int, status string) error {
	_, err := db.Conn.Exec(`
		UPDATE transactions SET status = $1, updated_at = NOW() WHERE id = $2
	`, status, transactionID)
	return err
}

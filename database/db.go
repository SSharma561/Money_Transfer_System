package database

import (
	"database/sql"
	"fmt"
	"money_transfer_system/config"

	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func Connect(cfg config.Config) (*Database, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Database{Conn: db}, nil
}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func (db *Database) UpdateBalances(senderID, receiverID int, amount float64) error {
	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	// Deduct amount from sender
	_, err = tx.Exec("UPDATE clients SET balance = balance - $1 WHERE id = $2 AND balance >= $1", amount, senderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Add amount to receiver
	_, err = tx.Exec("UPDATE clients SET balance = balance + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

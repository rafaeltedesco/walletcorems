package database

import (
	"database/sql"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{
		DB: db,
	}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := t.DB.Prepare(`
		INSERT INTO transactions (id, from_id, to_id, amount, created_at)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountFrom.ID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

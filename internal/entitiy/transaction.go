package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          string
	accountFrom *Account
	accountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		accountFrom: accountFrom,
		accountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	if err := transaction.Validate(); err != nil {
		return nil, err
	}
	transaction.Commit()
	return transaction, nil
}

func (t *Transaction) Commit() {
	t.accountFrom.Debit(t.Amount)
	t.accountTo.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greather than zero")
	}
	if t.accountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}

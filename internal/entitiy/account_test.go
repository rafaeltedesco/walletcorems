package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	account := NewAccount(customer)
	assert.NotNil(t, account)
	assert.Equal(t, customer.ID, account.Customer.ID)
}

func TestCreateAccountWithNilCustomer(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	account := NewAccount(customer)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	account := NewAccount(customer)
	account.Credit(100)
	account.Debit(45)
	assert.Equal(t, float64(55), account.Balance)
}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	customer1, _ := NewCustomer("Rafael", "rafael@gmail.com")
	accountRafael := NewAccount(customer1)
	customer2, _ := NewCustomer("Miguel", "miguel@gmail.com")
	accountMiguel := NewAccount(customer2)

	accountRafael.Credit(500)
	accountMiguel.Credit(200)

	transaction, err := NewTransaction(accountRafael, accountMiguel, 500)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 0.0, accountRafael.Balance)
	assert.Equal(t, 700.0, accountMiguel.Balance)
}

func TestCreateTransactionWithInsufficientFunds(t *testing.T) {
	customer1, _ := NewCustomer("Rafael", "rafael@gmail.com")
	accountRafael := NewAccount(customer1)
	customer2, _ := NewCustomer("Miguel", "miguel@gmail.com")
	accountMiguel := NewAccount(customer2)

	accountRafael.Credit(500)
	accountMiguel.Credit(200)

	transaction, err := NewTransaction(accountRafael, accountMiguel, 800)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())

}

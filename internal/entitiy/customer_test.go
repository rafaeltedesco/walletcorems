package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustommer(t *testing.T) {
	customer, err := NewCustomer("Rafael", "rafael@gmail.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Rafael", customer.Name)
	assert.Equal(t, "rafael@gmail.com", customer.Email)
}

func TestCreateNewCustomerWithInvalidArgs(t *testing.T) {
	customer, err := NewCustomer("", "")
	assert.NotNil(t, err)
	assert.Nil(t, customer)
}

func TestUpdateCustomer(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	err := customer.Update("Rafael Tedesco", "rafael_td@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "Rafael Tedesco", customer.Name)
	assert.Equal(t, "rafael_td@gmail.com", customer.Email)
}

func TestUpdateCustomerWithInvalidArgs(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	err := customer.Update("", "rafael_td@gmail.com")
	assert.NotNil(t, err)
	assert.Equal(t, "Name is required!", err.Error())
}

func TestAddAccount(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	account := NewAccount(customer)
	err := customer.AddAccount(account)
	assert.Nil(t, err)
	assert.Len(t, customer.Accounts, 1)
}

func TestAddInvalidAccount(t *testing.T) {
	customer, _ := NewCustomer("Rafael", "rafael@gmail.com")
	otherCustomer, _ := NewCustomer("Miguel", "miguel@gmail.com")
	otherAccount := NewAccount(otherCustomer)
	err := customer.AddAccount(otherAccount)
	assert.NotNil(t, err)
	assert.Equal(t, "account doesn't belong to these customer", err.Error())
}

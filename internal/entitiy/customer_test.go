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

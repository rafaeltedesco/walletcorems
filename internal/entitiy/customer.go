package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCustomer(name, email string) (*Customer, error) {
	customer := &Customer{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := customer.validate(); err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *Customer) validate() error {
	if c.Name == "" {
		return errors.New("Name is required!")
	}
	if c.Email == "" {
		return errors.New("Email is required!")
	}
	return nil
}

func (c *Customer) Update(name, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

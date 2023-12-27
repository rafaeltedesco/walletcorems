package createaccount

import (
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerGatewayMock struct {
	mock.Mock
}

func (m *CustomerGatewayMock) Save(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *CustomerGatewayMock) Get(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	customer, _ := entity.NewCustomer("Rafael", "rafael@gmail.com")

	c := &CustomerGatewayMock{}
	c.On("Get", customer.ID).Return(customer, nil)

	a := &AccountGatewayMock{}
	a.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(a, c)
	input := CreateAccountInputDTO{
		CustomerID: customer.ID,
	}
	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
	c.AssertExpectations(t)
	c.AssertNumberOfCalls(t, "Get", 1)
	a.AssertExpectations(t)
	a.AssertNumberOfCalls(t, "Save", 1)
}

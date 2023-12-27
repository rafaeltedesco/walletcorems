package createcustomer

import (
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CustomerGatewayMock struct {
	mock.Mock
}

func (m *CustomerGatewayMock) Get(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Customer), args.Error(1)
}

func (m *CustomerGatewayMock) Save(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func TestCreateCustomerUseCase_Execute(t *testing.T) {
	m := &CustomerGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateCustomerUseCase(m)
	output, err := uc.Execute(CreateCustomerInputDTO{Name: "Rafael", Email: "rafael@gmail.com"})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Rafael", output.Name)
	assert.Equal(t, "rafael@gmail.com", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}

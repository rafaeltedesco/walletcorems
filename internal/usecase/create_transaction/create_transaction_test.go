package createtransaction

import (
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	customer1, _ := entity.NewCustomer("rafael", "rafael@gmail.com")
	account1 := entity.NewAccount(customer1)
	account1.Credit(100)

	customer2, _ := entity.NewCustomer("miguel", "miguel@gmail.com")
	account2 := entity.NewAccount(customer2)
	account2.Credit(500)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindById", account1.ID).Return(account1, nil)
	mockAccount.On("FindById", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)
	output, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)

}

package createaccount

import (
	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com.br/rafaeltedesco/fc-walletcore/internal/gateway"
)

type CreateAccountInputDTO struct {
	CustomerID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway  gateway.AccountGateway
	CustomerGateway gateway.CustomerGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.CustomerGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway:  a,
		CustomerGateway: c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	customer, err := uc.CustomerGateway.Get(input.CustomerID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(customer)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}

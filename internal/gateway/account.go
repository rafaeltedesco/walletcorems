package gateway

import entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}

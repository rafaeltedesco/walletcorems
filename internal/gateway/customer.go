package gateway

import entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"

type CustomerGateway interface {
	Get(id string) (*entity.Customer, error)
	Save(customer *entity.Customer) error
}

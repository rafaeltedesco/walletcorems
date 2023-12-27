package gateway

import entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

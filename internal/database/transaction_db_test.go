package database

import (
	"database/sql"
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	transactionDB *TransactionDB
	customerFrom  *entity.Customer
	customerTo    *entity.Customer
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), from_id varchar(255), to_id varchar(255), amount int, created_at date)")
	s.transactionDB = NewTransactionDB(db)
	s.customerFrom, _ = entity.NewCustomer("Rafael", "rafael@gmail.com")
	s.accountFrom = entity.NewAccount(s.customerFrom)
	s.accountFrom.Credit(1000)
	s.customerTo, _ = entity.NewCustomer("Miguel", "miguel@gmail.com")
	s.accountTo = entity.NewAccount(s.customerTo)
	s.accountTo.Credit(200)
}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE transactions")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE customers")
}

func TestTransactionDBSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 300.0)
	s.Nil(err)
	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}

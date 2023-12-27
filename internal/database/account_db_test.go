package database

import (
	"database/sql"
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	customer  *entity.Customer
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), customer_id varchar(255), balance int, created_at date)")
	s.accountDB = NewAccountDB(db)
	customer, _ := entity.NewCustomer("Rafael", "rafael@gmail.com")
	s.customer = customer
}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindById() {
	s.db.Exec(`
		INSERT INTO customers (id, name, email, created_at)
		VALUES (?, ?, ?, ?)
	`, s.customer.ID, s.customer.Name, s.customer.Email, s.customer.CreatedAt)
	account := entity.NewAccount(s.customer)
	err := s.accountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.accountDB.FindById(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Customer.ID, accountDB.Customer.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.Equal(account.Customer.ID, accountDB.Customer.ID)
	s.Equal(account.Customer.Name, accountDB.Customer.Name)
	s.Equal(account.Customer.Email, accountDB.Customer.Email)
}

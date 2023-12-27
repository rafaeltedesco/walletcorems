package database

import (
	"database/sql"
	"testing"

	entity "github.com.br/rafaeltedesco/fc-walletcore/internal/entitiy"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CustomerDBTestSuite struct {
	suite.Suite
	db         *sql.DB
	customerDB *CustomerDB
}

func (s *CustomerDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE customers (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.customerDB = NewCustomerDB(db)
}

func (s *CustomerDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE customers")
}

func TestCustomerDBTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerDBTestSuite))
}

func (s *CustomerDBTestSuite) TestGet() {
	customer, _ := entity.NewCustomer("Rafael", "rafael@gmail.com")
	s.customerDB.Save(customer)

	customerDB, err := s.customerDB.Get(customer.ID)
	s.Nil(err)
	s.Equal(customer.ID, customerDB.ID)
	s.Equal(customer.Name, customerDB.Name)
	s.Equal(customer.Email, customerDB.Email)

}

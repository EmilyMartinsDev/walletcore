package database

import (
	"database/sql"
	"testing"
	"wallet/internal/enitity"
	_"github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)
type SuiteTransactionDBTest struct{
	suite.Suite
	db *sql.DB
	client1 *entity.Client
	client2 *entity.Client
	accountFrom *entity.Account
	accountTo *entity.Account
	transactionDB *TransactionDB
}

func (suite * SuiteTransactionDBTest) SetupSuite(){
	db, err:= sql.Open("sqlite3",":memory:" )
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_from_id varchar(255), account_to_id varchar(255), amount int, created_at date)")

	client1, err:= entity.NewClient("Emilt", "teste@123")
	suite.Nil(err)
	suite.client1 = client1
	
	client2, err:= entity.NewClient("Emilt2", "teste2@123")
	suite.Nil(err)
	suite.client2 = client2

	accountFrom := entity.NewAccount(suite.client1)
	accountFrom.Balance = 1400
	suite.accountFrom = accountFrom

	accountTo := entity.NewAccount(suite.client2)
	accountTo.Balance = 100
	suite.accountTo = accountTo

	suite.transactionDB = NewTransaction(db)
}
func (suite *SuiteTransactionDBTest) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE transaction")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTransactionDBTest))
}

func (suite *SuiteTransactionDBTest) TestCreate() {
	transaction, err := entity.NewTransaction(100, suite.accountFrom, suite.accountTo )
	suite.Nil(err)
	err = suite.transactionDB.Create(transaction)
	suite.Nil(err)
}
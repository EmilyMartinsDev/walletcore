package database

import (
	"database/sql"
	"testing"
	"wallet/internal/enitity"
	_"github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)
type ClientDBTestSuite struct {
	suite.Suite
	db *sql.DB
	clientDB *ClientDb
}

func (suite *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	suite.clientDB = NewClientDB(db)
}
func (suite *ClientDBTestSuite) TearDownSuite() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client, _ := entity.NewClient("Emily", "emily@teste.com")
	err,_ := suite.clientDB.Save(client)
	suite.NotNil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("Emily", "Emily@teste.com")
	 suite.clientDB.Save(client)

	clientDB, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, clientDB.ID)
	suite.Equal(client.Name, clientDB.Name)
	suite.Equal(client.Email, clientDB.Email)
}
package entity
import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestNewCreateTransaction( t *testing.T){
	client, _:= NewClient("emily", "teste@123.com")
	client2, _:= NewClient("emily2", "teste2@123.com")
	accountFrom := NewAccount(client)
	accountFrom.Credit(10)
	accountTo := NewAccount(client2)
	 tr, _:= NewTransaction(4, accountFrom, accountTo)
	 assert.Equal(t,4.0, tr.AccountTo.Balance)
	 assert.Equal(t,6.0, tr.AccountFrom.Balance)
}
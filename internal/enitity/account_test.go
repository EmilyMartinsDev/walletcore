package entity
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountValid(t *testing.T){
	client, _:= NewClient("emily", "teste@123.com")

	account:= NewAccount(client)

	assert.Equal(t, "emily", account.Client.Name)
	assert.NotNil(t, account)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreditAndDebit(t *testing.T){
	client, _:= NewClient("emily", "teste@123.com")

	account:= NewAccount(client)
	account.Credit(10)
	assert.Equal(t, "emily", account.Client.Name)
	assert.NotNil(t, account)
	assert.Equal(t, 10.0, account.Balance)
	account.Debit(5)
	assert.Equal(t, 5.0, account.Balance)
}

func TestDebitWithError(t *testing.T) {
    client, _ := NewClient("emily", "teste@123.com")
    account := NewAccount(client)

    _, err := account.Debit(100)
    assert.Error(t, err)
}


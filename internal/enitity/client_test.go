package entity
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewCreateClient( t *testing.T){
	client, err:= NewClient("emily","teste@teste.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "emily", client.Name)
	assert.Equal(t, "teste@teste.com", client.Email)

}

func TestNewCreateClientWhenIsInvalid( t *testing.T){
	client, err:= NewClient("","teste@teste.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T){
	client, _ := NewClient("emily", "teste@teste.com")
	err := client.UpdateClient("emily_update", "teste@teste2.com")
	assert.Equal(t, "emily_update", client.Name)
	assert.Equal(t, "teste@teste2.com", client.Email)
	assert.Nil(t, err)
}

func TestUpdateClientWhenIsInvalid(t *testing.T){
	client, _ := NewClient("emily", "teste@teste.com")
	err := client.UpdateClient("", "teste@teste2.com")
	assert.NotNil(t, err)
}
func TestaddAccount(t *testing.T){
	client, _ := NewClient("emily", "teste@teste.com")
	account := NewAccount(client)
	result := client.addAccount(account)
	assert.Nil(t, result)
	assert.Equal(t, 1, len(client.accounts))
}
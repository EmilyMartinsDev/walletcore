package client_usecase

import(
	"testing"
	"wallet/internal/gateway/mocks"
	"github.com/stretchr/testify/assert"
)


func TesteCreateClientUseCase_Execute(t *testing.T){
	m := &mocks.MockClientGateway{}
	usecase:= ConstructorClientUseCase(m)
	output, err := usecase.execute(CreateClientInputDto{
		Name: "Emily",
		Email:"teste@123.com",
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Emily",output.Name)

}
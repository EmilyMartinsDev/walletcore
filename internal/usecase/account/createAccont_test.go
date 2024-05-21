package account_usecase

import (
	"testing"
	"wallet/internal/gateway/mocks"
	"wallet/internal/enitity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	clientGatewayMock := mocks.NewClientGatewayMock()
	clientGatewayMock.On("Get", client.ID).Return(client, nil)
	Account := entity.NewAccount(client)
	accountGatewayMock := mocks.NewAccountGatewayMock()
	accountGatewayMock.On("Save", mock.Anything).Return(Account, nil)

	useCase := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := useCase.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)

	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
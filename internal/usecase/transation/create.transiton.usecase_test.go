package createtransaction
import(
	"testing"
	"wallet/internal/gateway/mocks"
	"github.com/stretchr/testify/assert"
	"wallet/internal/enitity"
	"github.com/stretchr/testify/mock"
)

func  TestUseCaseTransaction (t *testing.T){
	client, _ := entity.NewClient("John Doe", "john@doe.com")
	client2, _ := entity.NewClient("John2 Doe2", "john@doe2.com")
	AccountFrom := entity.NewAccount(client)
	AccountTO := entity.NewAccount(client2)
	AccountFrom.Credit(100)
	AccountTO.Credit(20)
	mocksAccount := mocks.NewAccountGatewayMock()
	mocksAccount.On("Find", AccountFrom.ID).Return(AccountFrom, nil)
	mocksAccount.On("Find", AccountTO.ID).Return(AccountTO, nil)
	mockTransaction := mocks.NewTransactionMockRepository()
	
	mockTransaction.On("Create", mock.Anything).Return( nil)
	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: AccountFrom.ID,
		AccountIDTo: AccountTO.ID,
		Amount: 100,
	}

	useCase := NewCreateTransactionUseCase(mockTransaction, mocksAccount)
	output, err := useCase.Execute(inputDto)
	
	assert.Nil(t, err)
	assert.NotNil(t, output)

	mocksAccount.AssertExpectations(t)
	mocksAccount.AssertNumberOfCalls(t, "Find", 2)

	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
package mocks
import (

	"wallet/internal/enitity"

	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct{
	mock.Mock
}

func NewTransactionMockRepository() (*TransactionRepository){
	return &TransactionRepository{}
}

func(m *TransactionRepository) Create(transaction *entity.Transaction) ( error){
	args:= m.Called(transaction)
	return  args.Error(0)
}
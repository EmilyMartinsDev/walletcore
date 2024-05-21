package mocks
import (

	"wallet/internal/enitity"

	"github.com/stretchr/testify/mock"
)
type AccountGatewayMock struct {
	mock.Mock
}

func NewAccountGatewayMock() *AccountGatewayMock {
	return &AccountGatewayMock{}
}

func (m *AccountGatewayMock) Save(account *entity.Account) (*entity.Account, error) {
	args := m.Called(account)
	return args.Get(0).( *entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Find(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}
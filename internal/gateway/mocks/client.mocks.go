package mocks
import (

	"wallet/internal/enitity"

	"github.com/stretchr/testify/mock"
)


type MockClientGateway struct{
	mock.Mock
}
func NewClientGatewayMock() *MockClientGateway {
	return &MockClientGateway{}
}
func (m *MockClientGateway) Get(id string) (*entity.Client, error){
	args:= m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}
func (m *MockClientGateway) Save(client *entity.Client) (*entity.Client, error) {
	args := m.Called(client)
	return args.Get(0).(*entity.Client), args.Error(1)
}
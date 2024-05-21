package account_usecase
import (
	
	"wallet/internal/gateway"
	"wallet/internal/enitity"
)
type CreateAccountInputDto struct{
	ClientID string

}
type CreateAccountOutputDto  struct{
	ID string
}

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway: clientGateway,
	}
}

func (useCase *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := useCase.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
    
	account :=  entity.NewAccount(client)
	account, _ = useCase.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
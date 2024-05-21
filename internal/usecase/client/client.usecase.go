package client_usecase
import (
	"time"
	
	"wallet/internal/gateway"
	"wallet/internal/enitity"
)
type CreateClientInputDto struct{
	Name string
	Email string

}
type CreateClientOutputDto struct{
	ID string
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time

}

type CreateClientUseCase struct {
 gateway gateway.ClientGateway
}

func ConstructorClientUseCase(gateway gateway.ClientGateway) (*CreateClientUseCase){
	usecase:= &CreateClientUseCase{
		gateway: gateway,

	}

	return usecase
}

func (usecase * CreateClientUseCase) execute(input CreateClientInputDto) (CreateClientOutputDto, error){
	client, err:= entity.NewClient(input.Name, input.Email)
	if(err != nil){
		return CreateClientOutputDto{}, err
	}
	clientCreated, err := usecase.gateway.Save(client)
	if err != nil{
		return CreateClientOutputDto{}, err
	}

	return CreateClientOutputDto{
		ID: clientCreated.ID,
		Email: clientCreated.Email,
		Name: clientCreated.Name,
		CreatedAt : clientCreated.CreatedAt,
		UpdatedAt: clientCreated.UpdatedAt,
	}, nil;
}
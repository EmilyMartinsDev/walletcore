package gateway
import "wallet/internal/enitity"

type ClientGateway interface{
	Save(client *entity.Client) (*entity.Client,  error)
	Get(id string) (*entity.Client,  error)
} 
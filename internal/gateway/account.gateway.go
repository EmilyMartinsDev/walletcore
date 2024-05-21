package gateway
import "wallet/internal/enitity"

type AccountGateway interface{
	Save(account *entity.Account) (*entity.Account,  error)
	Find(id string) (*entity.Account,  error)
	
} 
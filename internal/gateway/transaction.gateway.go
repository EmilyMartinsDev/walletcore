package gateway
import "wallet/internal/enitity"
type TransactionGateway interface{
	Create(transaction *entity.Transaction) (error)
}
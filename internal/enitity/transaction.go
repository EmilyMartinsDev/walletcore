package entity
import(
	"errors"
	"time"
	"github.com/google/uuid"
)

type Transaction struct{
	ID string
	Amount float64
    AccountFrom *Account
	AccountTo *Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransaction(amount float64, AccountFrom *Account, AccountTo *Account) (*Transaction, error){
	transaction := &Transaction{
		ID: uuid.New().String(),
		Amount: amount,
		AccountFrom: AccountFrom,
		AccountTo: AccountTo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	error:= transaction.Validade()
	if (error != nil){
		return nil, error 
	}
	transaction.Commit()
	return transaction, nil
	
}

func (t *Transaction) Commit() (*Transaction, error) {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
	
	return t, nil

}

func (t *Transaction) Validade() error{
	if(t.AccountFrom == nil || t.AccountTo == nil){
		return errors.New("account is invalid")
	}
	if(t.Amount <= 0){
		return errors.New("ammount is invalid")
	}
	if(t.AccountFrom.Balance <= 0){
		return errors.New("balance is invalid")
	}
	return nil
}
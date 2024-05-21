package database
import (
	"database/sql"
	 "wallet/internal/enitity"
)

type TransactionDB struct{
	DB *sql.DB
}

func NewTransaction(db *sql.DB) (*TransactionDB){
    return &TransactionDB{
		DB: db,
	}

}

func (t *TransactionDB) Create(transaction *entity.Transaction) error{
	stmt, err:= t.DB.Prepare("INSERT INTO transactions (id, account_from_id, account_to_id, amount, created_at) values (?,?,?,?,?) ")

	if err != nil{
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.CreatedAt, transaction.Amount)
	if err != nil{
		return err
	}
	return nil
}
package entity
import ("time"
"errors"
"github.com/google/uuid")
type Client struct{
	ID string
	Name string
	Email string
	accounts [] *Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string )(*Client, error){
	client:= &Client{
		ID: uuid.New().String(),
		Name: name,
		Email: email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := client.Validade()
	if err != nil{
		return nil, err
	}
	return client,nil
}

func (c *Client ) UpdateClient(name string, email string) error{
	c.Name = name
	c.Email = email
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	err:= c.Validade()
	if err != nil{
		return err
	}
	return nil
}	

func (c *Client) Validade() error{
	if(c.Name == ""){
		return errors.New("name is required")
	}
	if(c.Email== ""){
		return errors.New("email is required")
	}
	return nil
}

func (c *Client ) addAccount(account *Account) error {
	if(account.Client.ID != c.ID){
		return errors.New("this account has a client")
	}
	c.accounts = append(c.accounts, account)
	return nil
}
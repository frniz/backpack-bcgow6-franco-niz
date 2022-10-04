package internal

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Emitter  string  `json:"emitter" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
	LastID() (int, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Transaction, error) {
	return transactions, nil
}

func (r *repository) Store(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {
	t := Transaction{id, code, currency, price, emmiter, receiver, date}
	transactions = append(transactions, t)
	lastID = id
	return t, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func NewRepository() Repository {
	return &repository{}
}

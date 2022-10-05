package internal

import "fmt"

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
	Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
	PartialUpdate(id int, code string, price float64) (Transaction, error)
	Delete(id int) error
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

func (r *repository) Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {

	t := Transaction{id, code, currency, price, emmiter, receiver, date}
	updated := false

	for i := range transactions {
		if transactions[i].ID == t.ID {
			transactions[i] = t
			updated = true
			break
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("No se encontro la transaccion con ID: %d", t.ID)
	}

	return t, nil
}

func (r *repository) PartialUpdate(id int, code string, price float64) (Transaction, error) {
	var t Transaction
	updated := false

	for i := range transactions {
		if transactions[i].ID == id {
			transactions[i].Code = code
			transactions[i].Price = price
			t = transactions[i]
			updated = true
			break
		}
	}

	if !updated {
		return Transaction{}, fmt.Errorf("No se encontro la transaccion con ID: %d", id)
	}

	return t, nil
}

func (r *repository) Delete(id int) error {
	deleted := false

	var index int

	for i := range transactions {
		if transactions[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("No se encontró la transacción con ID: %d", id)
	}

	transactions = append(transactions[:index], transactions[index+1:]...)
	return nil
}

func NewRepository() Repository {
	return &repository{}
}

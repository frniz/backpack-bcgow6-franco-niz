package internal

import (
	"backpack-bcgow6-franco-niz/testing/practica2/TT/ej1/pkg/store"
	"fmt"
)

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Emitter  string  `json:"emitter" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
	Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error)
	PartialUpdate(id int, code string, price float64) (Transaction, error)
	Delete(id int) error
	LastID() (int, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Transaction, error) {
	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return []Transaction{}, err
	}
	return ts, nil
}

func (r *repository) Store(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {
	var ts []Transaction
	err := r.db.Read(&ts)
	if err != nil {
		return Transaction{}, err
	}
	t := Transaction{id, code, currency, price, emmiter, receiver, date}
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return Transaction{}, err
	}

	return t, nil
}

func (r *repository) LastID() (int, error) {
	var ts []Transaction

	if err := r.db.Read(&ts); err != nil {
		return 0, err
	}

	if len(ts) == 0 {
		return 0, nil
	}

	return ts[len(ts)-1].ID, nil
}

func (r *repository) Update(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {

	t := Transaction{id, code, currency, price, emmiter, receiver, date}
	updated := false

	ts, err := r.GetAll()
	if err != nil {
		return Transaction{}, fmt.Errorf("Hubo un error con la base de datos.")
	}

	for i := range ts {
		if ts[i].ID == t.ID {
			ts[i] = t
			if err := r.db.Write(ts); err != nil {
				return Transaction{}, err
			}
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

	ts, err := r.GetAll()
	if err != nil {
		return Transaction{}, fmt.Errorf("Hubo un error con la base de datos.")
	}

	for i := range ts {
		if ts[i].ID == id {
			ts[i].Code = code
			ts[i].Price = price
			t = ts[i]
			if err := r.db.Write(ts); err != nil {
				return Transaction{}, err
			}
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

	var ts []Transaction
	if err := r.db.Read(&ts); err != nil {
		return err
	}

	for i := range ts {
		if ts[i].ID == id {
			index = i
			deleted = true
			break
		}
	}

	if !deleted {
		return fmt.Errorf("No se encontró la transacción con ID: %d", id)
	}

	ts = append(ts[:index], ts[index+1:]...)
	if err := r.db.Write(ts); err != nil {
		return err
	}

	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

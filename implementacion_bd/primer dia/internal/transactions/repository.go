package internal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

const (
	SAVE_TRANSACTION = "INSERT INTO transactions (id, code, currency, price, emitter, receiver, date) VALUES (?, ?, ?, ?, ?, ?, ?);"

	GET_ALL_TRANSACTIONS = ""

	GET_TRANSACTION = "SELECT * FROM transactions WHERE id=?;"

	GET_ONE_TRANSACTION = "SELECT * FROM transactions WHERE code=?"

	UPDATE_TRANSACTION = "UPDATE transactions SET code=?, currency=?, price=?, emitter=?, receiver=?, date=? WHERE id=?;"

	DELETE_TRANSACTION = ""

	EXIST_TRANSACTION = "SELECT m.id FROM movies m WHERE m.id=?"
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
	db *sql.DB
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_TRANSACTION, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetAll() ([]Transaction, error) {

	return nil, nil
}

func (r *repository) Store(id int, code, currency string, price float64, emmiter, receiver, date string) (Transaction, error) {
	return Transaction{}, nil
}

func (r *repository) LastID() (int, error) {
	return 0, nil
}

func (r *repository) Update(id int, code, currency string, price float64, emitter, receiver, date string) error {
	stm, err := r.db.Prepare(UPDATE_TRANSACTION)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(code, currency, price, emitter, receiver, date, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
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

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

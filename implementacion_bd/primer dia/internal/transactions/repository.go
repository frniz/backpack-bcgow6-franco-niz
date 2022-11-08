package internal

import (
	"context"
	"database/sql"
	"errors"
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
	Store(code, currency string, price float64, emitter, receiver, date string) (int64, error)
	Update(id int, code, currency string, price float64, emitter, receiver, date string) error
	PartialUpdate(id int, code string, price float64) (Transaction, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_TRANSACTION, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) Get(id int64) (Transaction, error) {
	row := r.db.QueryRow(GET_TRANSACTION, id)
	var transaction Transaction
	err := row.Scan(&transaction.ID, &transaction.Code, &transaction.Currency, &transaction.Price,
		&transaction.Emitter, &transaction.Receiver, &transaction.Date)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (r *repository) GetAll() ([]Transaction, error) {

	return nil, nil
}

func (r *repository) Store(code, currency string, price float64, emitter, receiver, date string) (int64, error) {
	stm, err := r.db.Prepare(SAVE_TRANSACTION) //preparamos la consulta
	if err != nil {
		return 0, err
	}
	defer stm.Close()

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(code, currency, price, emitter, receiver, date)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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

	return t, nil
}

func (r *repository) Delete(id int) error {

	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

package internal

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetOne_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expectedTransaction := Transaction{
		ID:       1,
		Code:     "MELI01",
		Currency: "Dolar",
		Price:    100.0,
		Emitter:  "MeLi",
		Receiver: "MeLi",
		Date:     "2022-11-11",
	}
	columns := []string{"id", "code", "currency", "price", "emitter", "receiver", "date"}
	expectedQuery := "SELECT id, code, currency, price, emitter, receiver, date FROM transactions"
	rows := sqlmock.NewRows(columns)
	rows.AddRow(expectedTransaction.ID, expectedTransaction.Code, expectedTransaction.Currency,
		expectedTransaction.Price, expectedTransaction.Emitter, expectedTransaction.Receiver, expectedTransaction.Date)
	mock.ExpectQuery(expectedQuery).WithArgs(expectedTransaction.Code).WillReturnRows(rows)
	repository := NewRepository(db)

	tr, err := repository.GetOne(expectedTransaction.Code)
	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, tr)
}

func Test_Store_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expectedQuery := "INSERT INTO transactions"
	mock.ExpectPrepare(expectedQuery)
	mock.ExpectExec(expectedQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	var transactionId int64 = 1
	repository := NewRepository(db)
	transaction := Transaction{
		ID: 1,
	}

	id, err := repository.Store(transaction.Code, transaction.Currency, transaction.Price,
		transaction.Emitter, transaction.Receiver, transaction.Date)
	assert.NoError(t, err)
	assert.Equal(t, transactionId, id)
}

func Test_Delete_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	transaction := Transaction{
		ID:       1,
		Code:     "MELI01",
		Currency: "Dolar",
		Price:    100.0,
		Emitter:  "MeLi",
		Receiver: "MeLi",
		Date:     "2022-11-11",
	}
	expectedQueryDelete := "DELETE FROM transactions"
	mock.ExpectPrepare(expectedQueryDelete)
	mock.ExpectExec(expectedQueryDelete).WillReturnResult(sqlmock.NewResult(0, 0))
	repository := NewRepository(db)

	err = repository.Delete(transaction.ID)
	assert.NoError(t, err)

}

func Test_Update_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	expectedQuery := "UPDATE transactions"
	mock.ExpectPrepare(expectedQuery)
	mock.ExpectExec(expectedQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepository(db)

	err = repository.Update(1, "", "", 0, "", "", "")
	assert.NoError(t, err)
}

package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (s *StubStore) Read(data interface{}) error {
	t1 := Transaction{
		ID:       1,
		Code:     "000",
		Currency: "CurrencyPrueba1",
		Price:    1.0,
		Emitter:  "Prueba1",
		Receiver: "Prueba1Receiver",
		Date:     "UnaFecha",
	}
	t2 := Transaction{
		ID:       2,
		Code:     "001",
		Currency: "CurrencyPrueba2",
		Price:    2.0,
		Emitter:  "Prueba2",
		Receiver: "Prueba2Receiver",
		Date:     "UnaFecha2",
	}

	ts := []Transaction{
		t1,
		t2,
	}

	dataStub := data.(*[]Transaction)
	*dataStub = ts

	return nil
}
func (s *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	ss := StubStore{}
	repo := NewRepository(&ss)
	esperado := []Transaction{
		{
			ID:       1,
			Code:     "000",
			Currency: "CurrencyPrueba1",
			Price:    1.0,
			Emitter:  "BeforeUpdate",
			Receiver: "Prueba1Receiver",
			Date:     "UnaFecha",
		},
		{
			ID:       2,
			Code:     "001",
			Currency: "CurrencyPrueba2",
			Price:    2.0,
			Emitter:  "Prueba2",
			Receiver: "Prueba2Receiver",
			Date:     "UnaFecha2",
		},
	}

	resultado, err := repo.GetAll()
	fmt.Println(resultado)

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)

}

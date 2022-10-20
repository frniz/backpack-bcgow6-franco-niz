package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ReadWasCalled bool
	TrBefore      Transaction
}

func (s *MockStore) Read(data interface{}) error {
	s.ReadWasCalled = true
	t1 := Transaction{
		ID:       1,
		Code:     "BeforeUpdate",
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
func (s *MockStore) Write(data interface{}) error {
	return nil
}

func TestPartialUpdate(t *testing.T) {

	ss := MockStore{
		ReadWasCalled: false,
		TrBefore: Transaction{
			ID:       1,
			Code:     "BeforeUpdate",
			Currency: "CurrencyPrueba1",
			Price:    1.0,
			Emitter:  "Prueba1",
			Receiver: "Prueba1Receiver",
			Date:     "UnaFecha",
		},
	}
	repo := NewRepository(&ss)
	esperado := Transaction{
		ID:       1,
		Code:     "After Update",
		Currency: "CurrencyPrueba1",
		Price:    1.0,
		Emitter:  "Prueba1",
		Receiver: "Prueba1Receiver",
		Date:     "UnaFecha",
	}

	resultado, err := repo.PartialUpdate(1, "After Update", 1.0)
	fmt.Println(resultado)

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)
	assert.NotEqual(t, ss.TrBefore, resultado)
	assert.True(t, ss.ReadWasCalled)

}

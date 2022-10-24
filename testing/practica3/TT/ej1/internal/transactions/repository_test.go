package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestUpdate(t *testing.T) {

	ss := MockStore{
		ReadWasCalled: false,
		TrBefore: Transaction{
			ID:       1,
			Code:     "BeforeUpdate",
			Currency: "CurrencyPruebaBefore",
			Price:    1.0,
			Emitter:  "PruebaBefore",
			Receiver: "PruebaBefore",
			Date:     "UnaFechaBefore",
		},
	}
	repo := NewRepository(&ss)
	esperado := Transaction{
		ID:       1,
		Code:     "AfterUpdate",
		Currency: "CurrencyPruebaAfter",
		Price:    1.0,
		Emitter:  "PruebaAfter",
		Receiver: "PruebaAfter",
		Date:     "UnaFechaAfter",
	}

	resultado, err := repo.Update(1, "AfterUpdate", "CurrencyPruebaAfter", 1.0, "PruebaAfter", "PruebaAfter", "UnaFechaAfter")
	fmt.Println(resultado)

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)
	assert.NotEqual(t, ss.TrBefore, resultado)
	assert.True(t, ss.ReadWasCalled)

}

func TestLastID(t *testing.T) {
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

	beforeDB := []Transaction{
		t1,
		t2,
	}

	ss := MockStore{
		db: beforeDB,
	}
	repo := NewRepository(&ss)
	esperado := 2

	resultado, err := repo.LastID()

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)
}

func TestStore(t *testing.T) {
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

	beforeDB := []Transaction{
		t1,
		t2,
	}

	ss := MockStore{
		db: beforeDB,
	}
	repo := NewRepository(&ss)
	esperado := Transaction{
		ID:       3,
		Code:     "AfterUpdate",
		Currency: "CurrencyPruebaAfter",
		Price:    1.0,
		Emitter:  "PruebaAfter",
		Receiver: "PruebaAfter",
		Date:     "UnaFechaAfter",
	}

	resultado1, err1 := repo.Store(esperado.ID, esperado.Code, esperado.Currency, esperado.Price, esperado.Emitter,
		esperado.Receiver, esperado.Date)
	resultado2, err2 := repo.GetAll()

	assert.Nil(t, err1)
	assert.Equal(t, esperado, resultado1)
	assert.Nil(t, err2)
	assert.Equal(t, esperado, resultado2[len(resultado2)-1])
}

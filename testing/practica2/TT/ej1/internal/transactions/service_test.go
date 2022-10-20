package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialUpdateService(t *testing.T) {
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
	service := NewService(repo)
	esperado := Transaction{
		ID:       1,
		Code:     "After Update",
		Currency: "CurrencyPrueba1",
		Price:    2.0,
		Emitter:  "Prueba1",
		Receiver: "Prueba1Receiver",
		Date:     "UnaFecha",
	}

	resultado, err := service.PartialUpdate(ss.TrBefore.ID, "After Update", 2.0)

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)
	assert.True(t, ss.ReadWasCalled)
}

func TestDeleteService(t *testing.T) {
	idBorrar1 := 2
	idBorrar2 := 3
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
	afterDB := []Transaction{
		t1,
	}
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
		db:            beforeDB,
		errorEsperado: fmt.Errorf("No se encontró la transacción con ID: %d", idBorrar2),
	}
	repo := NewRepository(&ss)
	service := NewService(repo)

	err1 := service.Delete(idBorrar1)
	err2 := service.Delete(idBorrar2)

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
	assert.Equal(t, err2, ss.errorEsperado)
	assert.Equal(t, afterDB, ss.db)
}

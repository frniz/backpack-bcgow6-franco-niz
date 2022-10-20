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

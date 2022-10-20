package ej1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuma(t *testing.T) {
	num1 := 1
	num2 := 1
	resultadoEsperado := 2

	resultado := suma(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestResta(t *testing.T) {
	num1 := 1
	num2 := 1
	resultadoEsperado := 0

	resultado := resta(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestMult(t *testing.T) {
	num1 := 2
	num2 := 2
	resultadoEsperado := 4

	resultado := mult(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado)
}

func TestDiv(t *testing.T) {
	num1 := 2
	num2 := 2
	num3 := 0
	resultadoEsperado1 := 1

	resultado1, err1 := div(num1, num2)
	_, err2 := div(num1, num3)

	assert.Equal(t, resultadoEsperado1, resultado1)
	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}

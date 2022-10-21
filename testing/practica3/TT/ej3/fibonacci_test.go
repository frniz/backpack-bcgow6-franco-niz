package ej3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- FIBONACCI TDD ---
// caso large < 1: []
// Caso large = 1: [0]
// caso large = 2: [0, 1]
// caso large > 2: [0, 1, pos1 - pos0, ...]

func TestFibonacci(t *testing.T) {
	// a...
	// valorEsperado1 = nil
	valorEsperado2 := []int{0}
	valorEsperado3 := []int{0, 1}
	valorEsperado4 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

	parametro1 := -1
	parametro2 := len(valorEsperado2)
	parametro3 := len(valorEsperado3)
	parametro4 := len(valorEsperado4)

	// act
	resultado1 := Fibonacci(parametro1)
	resultado2 := Fibonacci(parametro2)
	resultado3 := Fibonacci(parametro3)
	resultado4 := Fibonacci(parametro4)

	//assert
	assert.Nil(t, resultado1)
	assert.Equal(t, valorEsperado2, resultado2)
	assert.Equal(t, valorEsperado3, resultado3)
	assert.Equal(t, valorEsperado4, resultado4)
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(2000000)
	}
}

package ej1

import "fmt"

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("El segundo valor no puede ser 0")
	}
	return a / b, nil
}

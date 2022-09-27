package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (prom int, err error) {

	if len(notas) == 0 {
		return 0, errors.New("Se ingresaron 0 notas")
	}

	for _, nota := range notas {

		if nota < 0 {
			return 0, errors.New("Hay una nota negativa")
		}

		prom += nota

	}

	return prom / len(notas), err
}

func main() {
	promedio, err := promedio(8, 8, 10, 7, 6, 5, 4)
	fmt.Printf("El promedio de notas es: %d", promedio)
	fmt.Printf("\n\n%s\n", err)
}

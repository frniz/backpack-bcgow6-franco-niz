package main

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones
// de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
// y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una
// cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(numbers ...int) (result int) {

	result = 99999
	for _, num := range numbers {
		if num < result {
			result = num
		}
	}

	return
}

func maxFunc(numbers ...int) (result int) {

	for _, num := range numbers {
		if num > result {
			result = num
		}
	}

	return
}

func averageFunc(numbers ...int) (result int) {

	for _, num := range numbers {
		result += num
	}

	return result / len(numbers)
}

func operation(operationToDo string) (func(...int) int, error) {

	switch operationToDo {

	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil

	}

	return nil, errors.New("Invalid operation.")

}

func main() {
	min, err := operation(minimum)
	fmt.Printf("\n%s", err)
	average, err := operation(average)
	fmt.Printf("\n%s", err)
	max, err := operation(maximum)
	fmt.Printf("\n%s", err)

	minValue := min(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := average(2, 3, 3, 4, 10, 2, 4, 5)
	maxValue := max(2, 3, 3, 4, 10, 2, 4, 5)

	fmt.Printf("\nMinimo de notas: %d\nPromedio de notas: %d\nMaximo de notas: %d\n", minValue, averageValue, maxValue)

	_, err = operation("indefined")
	fmt.Printf("\nIndefined operation: %s\n", err)
}

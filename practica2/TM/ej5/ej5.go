package main

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo
// tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

// - perro necesitan 10 kg de alimento
// - gato 5 kg
// - Hamster 250 gramos.
// - Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que
// retorne una función y un mensaje (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func alimentoPerro(cantidad int) float64 {
	return float64(cantidad) * 10
}

func alimentoGato(cantidad int) float64 {
	return float64(cantidad) * 5
}

func alimentoHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func alimentoTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

func Animal(tipo string) (func(int) float64, error) {

	switch tipo {
	case dog:
		return alimentoPerro, nil
	case cat:
		return alimentoGato, nil
	case hamster:
		return alimentoHamster, nil
	case tarantula:
		return alimentoTarantula, nil
	}

	return nil, errors.New("That animal is not defined")

}

func main() {
	animalDog, msg := Animal(dog)
	animalCat, msg := Animal(cat)

	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)

	fmt.Printf("%s", msg)

}

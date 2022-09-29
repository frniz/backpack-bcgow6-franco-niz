package main

// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
// represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
// si es cuadrática y cuál es el valor máximo.

import "fmt"

type Matrix struct {
	Matriz       [5][]float64
	Alto         int
	Ancho        int
	EsCuadratica bool
	ValorMaximo  float64
}

func (m *Matrix) Set(cantColum int, values ...float64) {

	m.Alto = 5
	m.Ancho = cantColum
	m.EsCuadratica = cantColum == 5

	for i := 0; i < m.Alto; i++ {

		m.Matriz[i] = make([]float64, cantColum)

	}

	var iColum, iFila int
	var maximo float64
	for _, v := range values {
		if iColum < cantColum {
			if v > maximo {
				maximo = v
			}

			m.Matriz[iFila][iColum] = v
			iColum++

			if (iColum % cantColum) == 0 {
				iColum = 0
				iFila++
			}
		}

	}

	m.ValorMaximo = maximo

}

func (m Matrix) Print() {

	for i := 0; i < m.Alto; i++ {
		fmt.Println(m.Matriz[i])
	}

}

func main() {

	matriz := Matrix{}
	matriz.Set(3, 0.5, 3, 22.4, 14.1, 11, 9.8, 6.4, 1, -1.5, 101.2, 48.3, 12, 95.4, 0.251)
	matriz.Print()

}

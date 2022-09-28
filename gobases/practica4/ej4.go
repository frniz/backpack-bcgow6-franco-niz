// Vamos a hacer que nuestro programa sea un poco más complejo.
// Desarrolla las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver
// un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
// Calcular el medio aguinaldo correspondiente al trabajador
// Fórmula de cálculo de aguinaldo:
// [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
// La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un
// número negativo.

// Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
// “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.

package main

import (
	"fmt"
)

const (
	MSJ_CERO = "error no identificado"
	MSJ_UNO  = "el trabajador no puede haber trabajado menos de 80 hs mensuales"
	MSJ_DOS  = "se ingreso un numero negativo"
)

type ErrorSalarios struct {
	ID int
}

func (e *ErrorSalarios) Error() string {
	msj := ""
	switch e.ID {
	case 1:
		msj = fmt.Sprint("Error: ", MSJ_UNO)

	case 2:
		msj = fmt.Sprint("Error: ", MSJ_DOS)

	default:
		msj = fmt.Sprint("Error: ", MSJ_CERO)
	}
	return msj
}

func SalarioMensual(horasTrabajadas int, valorDeLaHora float64) (salario float64, err error) {

	if horasTrabajadas < 80 {
		err = &ErrorSalarios{ID: 1}
		return
	}

	salario = float64(horasTrabajadas) * valorDeLaHora

	if salario > 150_000 {
		salario -= salario * 0.1
	}

	return
}

func MedioAguinaldo(mesesTrabajadosEnElSemestre int, salariosDelSemestre [6]float64) (medioAguinaldo float64, err error) {

	if mesesTrabajadosEnElSemestre < 0 {
		err = &ErrorSalarios{ID: 2}
		return
	}

	var maximoSalario float64
	for _, v := range salariosDelSemestre {
		if v > maximoSalario {
			maximoSalario = v
		} else if v < 0 {
			err = &ErrorSalarios{ID: 2}
			return
		}
	}

	medioAguinaldo = maximoSalario / 12 * float64(mesesTrabajadosEnElSemestre)
	return
}

func main() {

	salarios := [6]float64{150_000, 89_000, 120_000}

	fmt.Printf("\nIngrese la cantidad de horas trabajadas: ")
	var horas int
	fmt.Scan(&horas)
	fmt.Printf("\nIngrese el valor de la horas: ")
	var salarioHora float64
	fmt.Scan(&salarioHora)

	salarioMensual, err := SalarioMensual(horas, salarioHora)
	fmt.Printf("\n%f", salarioMensual)

	if err != nil {
		fmt.Printf("\nerr: %v", err.Error())
	}

	medioAguinaldo, err2 := MedioAguinaldo(4, salarios)
	fmt.Printf("\n%f", medioAguinaldo)

	if err2 != nil {
		fmt.Printf("\nerr2: %v\n", err2.Error())
	}
}

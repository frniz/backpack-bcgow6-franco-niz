// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba
// por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por
// consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”,
// siendo [salary] el valor de tipo int pasado por parámetro).

package main

import (
	"fmt"
	"os"
)

type CustomError struct {
}

// func (c *CustomError) Error() string {
// 	return "El salario ingresado no alcanza el minimo imponible"
// }

func main() {

	salary := 293_000.0

	if salary < 150_000 {
		// c := CustomError{}
		// fmt.Printf("c.Error(): %v\n", c.Error())

		// e := errors.New("El salario ingresado no alcanza el minimo imponible")
		// fmt.Printf("e: %v\n", e)

		e := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %.2f", salary)
		fmt.Printf("e: %v\n", e)

		os.Exit(0)
	}

	fmt.Println("Debe pagar impuestos.")

}

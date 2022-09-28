// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
// se implemente “errors.New()”.

package main

import (
	"errors"
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

		e := errors.New("El salario ingresado no alcanza el minimo imponible")
		fmt.Printf("e: %v\n", e)

		os.Exit(0)
	}

	fmt.Println("Debe pagar impuestos.")

}

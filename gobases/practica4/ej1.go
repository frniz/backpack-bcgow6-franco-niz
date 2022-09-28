// En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
// Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario
// ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario,
// imprime por consola el mensaje “Debe pagar impuesto”.

package main

import (
	"fmt"
	"os"
)

type CustomError struct {
}

func (c *CustomError) Error() string {
	return "El salario ingresado no alcanza el minimo imponible"
}

func main() {

	salary := 293_000.0

	if salary < 150_000 {
		c := CustomError{}
		fmt.Printf("c.Error(): %v\n", c.Error())
		os.Exit(0)
	}

	fmt.Println("Debe pagar impuestos.")

}

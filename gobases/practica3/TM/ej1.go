// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados,
// separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

package main

import (
	"fmt"
	"os"
)

func leerProducto() {

	f, err := os.OpenFile("productos.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	var entrada1 int
	var entrada2, entrada3, entrada4 string
	fmt.Println("Escriba el ID: ")
	fmt.Scan(&entrada1)

	if entrada1 == 0 {
		f.Close()
		os.Exit(0)
	}

	fmt.Println("Escriba el producto: ")
	fmt.Scan(&entrada2)
	fmt.Println("Escriba el precio: $")
	fmt.Scan(&entrada3)
	fmt.Println("Escriba la cantidad: ")
	fmt.Scan(&entrada4)

	producto := fmt.Sprint(entrada1, ",", entrada2, ",$", entrada3, ",", entrada4, ";")
	fmt.Fprint(f, producto)

}

func main() {

	fmt.Println("Ingreso de productos: (0 para terminar)")

	for true {
		leerProducto()
	}

}

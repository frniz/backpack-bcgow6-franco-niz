package main

import "fmt"

func impuesto(sueldo float64) (impuesto float64) {

	if sueldo > 150_000 {
		impuesto = 0.27 * sueldo
	} else if sueldo > 50_000 {
		impuesto = 0.17 * sueldo
	}

	return
}

func main() {
	fmt.Println("Ingrese el sueldo del empleado: ")

	var sueldo float64
	fmt.Scanln(&sueldo)

	fmt.Printf("\nEl impuesto a cobrarle es: %0.2f\n\n", impuesto(sueldo))
}

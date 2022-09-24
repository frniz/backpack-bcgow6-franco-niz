package main

import "fmt"

func main() {

	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio",
		7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	fmt.Println("Ingrese el numero de mes del cual quiera saber el nombre: ")

	var numeroDeMes int
	fmt.Scanln(&numeroDeMes)

	if numeroDeMes > 0 && numeroDeMes < 13 {
		fmt.Println("\n\n", meses[numeroDeMes])
	} else {
		fmt.Println("\n\nNo corresponde al numero de un mes")
	}

}

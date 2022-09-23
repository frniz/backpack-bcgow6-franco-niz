package main

/*
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	var count int

	for _, age := range employees {

		if age > 21 {
			count++
		}

	}

	fmt.Printf("Cantidad de empleados mayores a 21: %d\n", count)

	employees["Federico"] = 25
	delete(employees, "Pedro")

	for name, age := range employees {
		fmt.Printf("%s: %d\n", name, age)
	}

}

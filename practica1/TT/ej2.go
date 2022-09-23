package main

import "fmt"

func main() {

	edad := 20
	empleado := true
	antiguedadEnAnios := 1
	sueldo := 100000

	fmt.Println("Otorga credito?:")

	if edad < 22 {
		fmt.Println("No, ya que no tiene la edad minima")
	} else if empleado == false {
		fmt.Println("No, ya que no esta empleado")
	} else if antiguedadEnAnios < 1 {
		fmt.Println("No, ya que no tiene suficiente antigüedad")
	} else if sueldo < 100000 {
		fmt.Println("Si, y se le cobraran intereses")
	} else {
		fmt.Println("Si, y no se le cobraran intereses")
	}

}

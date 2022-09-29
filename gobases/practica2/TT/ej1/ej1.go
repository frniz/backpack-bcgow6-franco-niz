package main

import (
	"fmt"
)

type Fec struct {
	dia int
	mes int
	año int
}

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    Fec
}

func (f Fec) toString() string {
	return fmt.Sprint(f.dia) + "/" + fmt.Sprint(f.mes) + "/" + fmt.Sprint(f.año)
}

func (e Estudiante) detalle() string {
	return "Soy " + fmt.Sprint(e.Nombre) + " " + fmt.Sprint(e.Apellido) + ", DNI: " + fmt.Sprint(e.DNI) + " y naci el " +
		e.Fecha.toString()
}

func main() {
	alumno := Estudiante{
		Nombre:   "Franco",
		Apellido: "Niz",
		DNI:      43574,
		Fecha: Fec{
			dia: 9,
			mes: 20,
			año: 2001,
		},
	}

	fmt.Printf("\n%s\n", alumno.detalle())
}

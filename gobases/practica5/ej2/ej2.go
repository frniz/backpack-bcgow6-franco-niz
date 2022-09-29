package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Celular   string
	Domicilio string
}

func GenerarID() (result int) {

	result = rand.Intn(100)

	if result == 0 {
		panic("El ID generado es invalido.")
	}

	return
}

func ArchivoNoExiste() {
	err := recover()

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func ExisteCliente(c Cliente) {

	_, err := os.Open("customers.txt")
	defer ArchivoNoExiste()
	if err != nil {
		panic("El archivo no existe.")
	}

	// Aca se compararia el archivo con el cliente nuevo

}

func ComprobarCero(c Cliente) (bool, error) {

	if c.Apellido == "" || c.Nombre == "" || c.Celular == "" || c.Domicilio == "" || c.DNI == 0 {
		return true, errors.New("Existe un valor 0 en el cliente")
	}

	return false, nil
}

func main() {
	var cliente Cliente

	defer fmt.Println("No han quedado archivos abiertos")
	defer fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
	defer fmt.Println("Final de ejecucion")

	cliente.Legajo = GenerarID()

	fmt.Println("Ingrese el nombre")
	var nombre string
	fmt.Scanln(nombre)
	fmt.Println("Ingrese el apellido")
	fmt.Scan(cliente.Apellido)
	fmt.Println("Ingrese el domicilio")
	fmt.Scan(cliente.Domicilio)
	fmt.Println("Ingrese el celular")
	fmt.Scan(cliente.Celular)
	fmt.Println("Ingrese el DNI")
	fmt.Scan(cliente.DNI)

	_, err := ComprobarCero(cliente)
	defer ArchivoNoExiste()

	if err != nil {
		panic("Existe algun 0 en los datos")
	}

	ExisteCliente(cliente)

}

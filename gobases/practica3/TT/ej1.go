// Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando
// información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe
// el mismo lugar en memoria para el main del programa y para las funciones.
// La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
// Y deben implementarse las funciones:
// Cambiar nombre: me permite cambiar el nombre y apellido.
// Cambiar edad: me permite cambiar la edad.
// Cambiar correo: me permite cambiar el correo.
// Cambiar contraseña: me permite cambiar la contraseña.

package main

import (
	"fmt"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func CambiarNombre(u *Usuario, nombre string) {
	u.Nombre = nombre
}

func CambiarEdad(u *Usuario, edad int) {
	u.Edad = edad
}

func CambiarCorreo(u *Usuario, correo string) {
	u.Correo = correo
}

func cambiarConstraseña(u *Usuario, contraseña string) {
	u.Contraseña = contraseña
}

func main() {

	u := Usuario{
		Nombre:     "Franco",
		Apellido:   "Niz",
		Edad:       20,
		Correo:     "franco.niz@mercadolibre.com",
		Contraseña: "1234abcd",
	}

	fmt.Println(u)

	CambiarNombre(&u, "Damian")

	fmt.Println(u)

}

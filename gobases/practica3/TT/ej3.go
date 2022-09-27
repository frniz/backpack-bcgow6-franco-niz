// Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
// Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios
// y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la
// sumatoria se realice en paralelo mediante 3 go routines.

// Se requieren 3 estructuras:
// Productos: nombre, precio, cantidad.
// Servicios: nombre, precio, minutos trabajados.
// Mantenimiento: nombre, precio.

// Se requieren 3 funciones:
// Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
// Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
// si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
// Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

// Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final
// (sumando el total de los 3).

package main

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(productos []Productos, c chan float64) {

	precioTotal := 0.0
	for _, value := range productos {
		precioTotal += value.Precio * float64(value.Cantidad)
	}

	c <- precioTotal

}

func SumarServicios(servicios []Servicios, c chan float64) {

	precioTotal := 0.0
	for _, value := range servicios {
		mediaHoras := value.MinutosTrabajados / 30
		if mediaHoras == 0 {
			mediaHoras++
		}
		precioTotal += float64(mediaHoras) * value.Precio
	}

	c <- precioTotal

}

func SumarMantenimientos(mantenimientos []Mantenimiento, c chan float64) {

	precioTotal := 0.0
	for _, value := range mantenimientos {
		precioTotal += value.Precio
	}

	c <- precioTotal
}

func main() {

	// c1 := make(chan float64)
	// c2 := make(chan float64)
	// c3 := make(chan float64)
	// go SumarProductos(p, c1)
	// go SumarServicios(s, c2)
	// go SumarMantenimientos(m, c3)

	// v1 := <-c1
	// v2 := <-c2
	// v3 := <-c3

}

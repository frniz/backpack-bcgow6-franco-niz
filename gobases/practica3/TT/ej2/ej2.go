// Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
// Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main d
// el programa como en las funciones.
// Se necesitan las estructuras:
// Usuario: Nombre, Apellido, Correo, Productos (array de productos).
// Producto: Nombre, precio, cantidad.
// Se requieren las funciones:
// Nuevo producto: recibe nombre y precio, y retorna un producto.
// Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
// Borrar productos: recibe un usuario, borra los productos del usuario.

package main

type Producto struct {
	Nombre   string
	Precio   float64
	cantidad int
}

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

func NuevoProducto(nombre string, precio float64) (p Producto) {

	p.Nombre = nombre
	p.Precio = precio

	return

}

func AgregarProducto(u *Usuario, p Producto, c int) {
	p.cantidad = c
	u.Productos = append(u.Productos, p)
}

func BorrarProductos(u *Usuario) {
	u.Productos = []Producto{}
}

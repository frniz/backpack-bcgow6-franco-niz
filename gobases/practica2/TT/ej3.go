// Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
// retornar el valor del precio total.
// Las empresas tienen 3 tipos de productos:
// Pequeño, Mediano y Grande. (Se espera que sean muchos más)
// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

// Sus costos adicionales son:
// Pequeño: El costo del producto (sin costo adicional)
// Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
// Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

// Requerimientos:
// Crear una estructura “tienda” que guarde una lista de productos.
// Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
// Crear una interface “Producto” que tenga el método “CalcularCosto”
// Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
// Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
// Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
// Interface Producto:
// El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
// Interface Ecommerce:
//  - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si
//  los hubiera.
//  - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

package main

const (
	Grande  = "Grande"
	Mediano = "Mediano"
	Pequeño = "Pequeño"
)

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type tienda struct {
	productos []producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) (p producto) {
	p.nombre = nombre
	p.tipo = tipo
	p.precio = precio

	return
}

func nuevaTienda() (e Ecommerce) {
	return
}

func (t tienda) Total() (total float64) {
	for _, v := range t.productos {
		total += v.precio
	}

	return
}

func (t tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
}

func (p producto) CalcularCosto() (total float64) {
	switch p.tipo {
	case Grande:
		total = p.precio
	case Mediano:
		total = p.precio + p.precio*0.03
	case Pequeño:
		total = p.precio + p.precio*0.06 + 2500.0
	}

	return
}

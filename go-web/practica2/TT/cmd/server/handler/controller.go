package controller

// Se debe separar la estructura del proyecto, como segundo paso se debe generar el paquete server donde se agregaran las
// funcionalidades del proyecto que dependan de paquetes externos y el main del programa.

// Dentro del paquete deben estar:
// El main del programa.
// - Se debe importar e inyectar el repositorio, servicio y handler
// - Se debe implementar el router para los diferentes endpoints
// El paquete handler con el controlador de la entidad seleccionada.
// - Se debe generar la estructura request
// - Se debe generar la estructura del controlador que tenga como campo el servicio
// - Se debe generar la función que retorne el controlador
// - Se deben generar todos los métodos correspondientes a los endpoints

type request struct {
	ID       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Emitter  string  `json:"emitter" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

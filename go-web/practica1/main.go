package main

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

// Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
// Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
// Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
// Genera un handler para el endpoint llamado “GetAll”.
// Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

// Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
// Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
// Luego genera la lógica de filtrado de nuestro array.
// Devolver por el endpoint el array filtrado.

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

}

func main() {

	type Transactions struct {
		ID       int
		Code     string
		Currency string
		Price    float64
		Emitter  string
		Receiver string
		Date     string
	}

	content, err := os.ReadFile("./transactions.json")
	if err != nil {
		panic(err)
	}

	var transactions []Transactions
	err = json.Unmarshal(content, &transactions)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":      "Hola Franco!",
			"transactions": transactions,
		})
	})
	router.Run()
}

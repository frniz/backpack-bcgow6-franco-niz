package main

// Ejercicio 1 - Configuración ENV

// Configurar para que el token sea tomado de las variables de entorno al momento de realizar
// la validación, para eso se deben realizar los siguientes pasos:
// 1. Configurar la aplicación para que tome los valores que se encuentran en el archivo
// .env como variable de entorno.
// 2. Quitar el valor del token del código y agregar como variable de entorno.
// 3. Acceder al valor del token mediante la variable de entorno.

// Ejercicio 2 - Guardar información
// Se debe implementar la funcionalidad para guardar la información de la petición en un
// archivo json, para eso se deben realizar los siguientes pasos:
// 1. En lugar de guardar los valores de nuestra entidad en memoria, se debe crear un
// archivo; los valores que se vayan agregando se guardan en él.

// Ejercicio 3 - Leer información

// Se debe implementar la funcionalidad para leer la información requerida en la petición del
// archivo json generado al momento de guardar, para eso se deben realizar los siguientes
// pasos:

// 1. En lugar de leer los valores de nuestra entidad en memoria, se debe obtener del
// archivo generado en el punto anterior.

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	handler "backpack-bcgow6-franco-niz/go-web/practica4/TM/cmd/server/handler"
	transactions "backpack-bcgow6-franco-niz/go-web/practica4/TM/internal/transactions"
	"backpack-bcgow6-franco-niz/go-web/practica4/TM/pkg/store"
)

func main() {
	_ = godotenv.Load()
	db := store.New(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()
	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.PartialUpdate())
	tr.DELETE("/:id", t.Delete())
	r.Run()
}

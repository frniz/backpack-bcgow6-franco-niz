package main

// Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
// lograrlo, es necesario, seguir los siguientes pasos:
// 1. Generar un método PUT para modificar la entidad completa
// 2. Desde el Path enviar el ID de la entidad que se modificará
// 3. En caso de no existir, retornar un error 404
// 4. Realizar todas las validaciones (todos los campos son requeridos)

// Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
// deben modificar 2 campos:
// - Si se seleccionó Productos, los campos nombre y precio.
// - Si se seleccionó Usuarios, los campos apellido y edad.
// - Si se seleccionó Transacciones, los campos código de transacción y monto.
// .Para lograrlo, es necesario, seguir los siguientes pasos:
// 1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
// campo (a elección)
// 2. Desde el Path enviar el ID de la entidad que se modificara
// 3. En caso de no existir, retornar un error 404

// Es necesario implementar una funcionalidad para eliminar una entidad. Para lograrlo, es
// necesario, seguir los siguientes pasos:
// 1. Generar un método DELETE para eliminar la entidad en base al ID
// 2. En caso de no existir, retornar un error 404

import (
	"github.com/gin-gonic/gin"

	handler "backpack-bcgow6-franco-niz/go-web/practica3/TM/cmd/server/handler"
	transactions "backpack-bcgow6-franco-niz/go-web/practica3/TM/internal/transactions"
)

func main() {
	repo := transactions.NewRepository()
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

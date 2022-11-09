package main

import (
	"os"

	"github.com/joho/godotenv"

	handler "backpack-bcgow6-franco-niz/implementacion_bd/primer_dia/cmd/server/handler"
	"backpack-bcgow6-franco-niz/implementacion_bd/primer_dia/docs"
	transactions "backpack-bcgow6-franco-niz/implementacion_bd/primer_dia/internal/transactions"
	"backpack-bcgow6-franco-niz/implementacion_bd/primer_dia/pkg/store"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name Franco Niz
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	_ = godotenv.Load()
	r, db := store.ConnectDatabase()
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.PartialUpdate())
	tr.DELETE("/:id", t.Delete())
	r.Run()
}

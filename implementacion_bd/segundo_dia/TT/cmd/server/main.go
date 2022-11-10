package main

import (
	"os"

	"github.com/joho/godotenv"

	handler "backpack-bcgow6-franco-niz/implementacion_bd/segundo_dia/TT/cmd/server/handler"
	"backpack-bcgow6-franco-niz/implementacion_bd/segundo_dia/TT/docs"
	transactions "backpack-bcgow6-franco-niz/implementacion_bd/segundo_dia/TT/internal/transactions"
	"backpack-bcgow6-franco-niz/implementacion_bd/segundo_dia/TT/pkg/store"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	ENV_PATH = "../../.env"
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
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	r, db, err := store.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.GET("/:id", t.Get())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.PartialUpdate())
	tr.DELETE("/:id", t.Delete())
	r.Run()
}

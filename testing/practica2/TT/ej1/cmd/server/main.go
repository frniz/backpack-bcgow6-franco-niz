package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	handler "backpack-bcgow6-franco-niz/testing/practica2/TT/ej1/cmd/server/handler"
	"backpack-bcgow6-franco-niz/testing/practica2/TT/ej1/docs"
	transactions "backpack-bcgow6-franco-niz/testing/practica2/TT/ej1/internal/transactions"
	"backpack-bcgow6-franco-niz/testing/practica2/TT/ej1/pkg/store"

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
	db := store.New(store.FileType, "./transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.PartialUpdate())
	tr.DELETE("/:id", t.Delete())
	err := r.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

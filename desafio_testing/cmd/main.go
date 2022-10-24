package main

import (
	"backpack-bcgow6-franco-niz/desafio_testing/cmd/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	//r.Run(":18085")
	r.Run()
}

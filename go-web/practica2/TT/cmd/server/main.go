package server

import (
	"github.com/gin-gonic/gin"

	handler "backpack-bcgow6-franco-niz/go-web/practica2/TT/cmd/server/handler"
	transactions "backpack-bcgow6-franco-niz/go-web/practica2/TT/internal/transactions"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()
	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	r.Run()
}

package server

import "github.com/gin-gonic/gin"

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := handlder.NewTransaction(service)

	r := gin.Default()
	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
}

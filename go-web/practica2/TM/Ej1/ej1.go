package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Se debe implementar la funcionalidad para crear la entidad. pasa eso se deben seguir los
// siguientes pasos:
// 1. Crea un endpoint mediante POST el cual reciba la entidad.
// 2. Se debe tener un array de la entidad en memoria (a nivel global), en el cual se
// deberán ir guardando todas las peticiones que se vayan realizando.
// 3. Al momento de realizar la petición se debe generar el ID. Para generar el ID se debe
// buscar el ID del último registro generado, incrementarlo en 1 y asignarlo a nuestro
// nuevo registro (sin tener una variable de último ID a nivel global).

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	Emitter  string  `json:"emitter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

var Transactions []Transaction

func EntityReceiver(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("transaction: %v\n", transaction)
	transaction.ID = len(Transactions) + 1
	Transactions = append(Transactions, transaction)
	c.JSON(http.StatusAccepted, transaction)
	fmt.Printf("Transactions: %v\n", Transactions)

}

func main() {
	r := gin.Default()
	r.POST("/transactions", EntityReceiver)
	r.Run()
}

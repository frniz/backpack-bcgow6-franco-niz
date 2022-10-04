package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Se debe implementar las validaciones de los campos al momento de enviar la petición, para
// eso se deben seguir los siguientes pasos:
// 1. Se debe validar todos los campos enviados en la petición, todos los campos son
// requeridos
// 2. En caso que algún campo no esté completo se debe retornar un código de error 400
// con el mensaje “el campo %s es requerido”.
// (En %s debe ir el nombre del campo que no está completo).

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Emitter  string  `json:"emitter" binding:"required"`
	Receiver string  `json:"receiver" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

var Transactions []Transaction

func EntityReceiver(c *gin.Context) {
	var transaction Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		var message string
		if transaction.Code == "" {
			message = message + fmt.Sprintf("El campo Code es requerido. ")
		}
		if transaction.Currency == "" {
			message = message + fmt.Sprintf("El campo Currency es requerido. ")
		}
		if transaction.Price == 0 {
			message = message + fmt.Sprintf("El campo Price es requerido. ")
		}
		if transaction.Emitter == "" {
			message = message + fmt.Sprintf("El campo Emitter es requerido. ")
		}
		if transaction.Receiver == "" {
			message = message + fmt.Sprintf("El campo Receiver es requerido. ")
		}
		if transaction.Date == "" {
			message = message + fmt.Sprintf("El campo Date es requerido. ")
		}

		c.String(400, message)
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

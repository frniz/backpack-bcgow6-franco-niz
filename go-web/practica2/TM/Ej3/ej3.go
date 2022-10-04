package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Para agregar seguridad a la aplicación se debe enviar la petición con un token, para eso se
// deben seguir los siguientes pasos::

// 1. Al momento de enviar la petición se debe validar que un token sea enviado
// 2. Se debe validar ese token en nuestro código (el token puede estar hardcodeado).
// 3. En caso que el token enviado no sea correcto debemos retornar un error 401 y un
// mensaje que “no tiene permisos para realizar la petición solicitada”.

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
	token := c.GetHeader("token")
	if token != "newPackage" {
		c.JSON(401, gin.H{
			"error": "This token is not valid.",
		})
		return
	}

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

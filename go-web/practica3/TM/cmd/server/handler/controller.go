package handler

// Se debe separar la estructura del proyecto, como segundo paso se debe generar el paquete server donde se agregaran las
// funcionalidades del proyecto que dependan de paquetes externos y el main del programa.

// Dentro del paquete deben estar:
// El main del programa.
// - Se debe importar e inyectar el repositorio, servicio y handler
// - Se debe implementar el router para los diferentes endpoints
// El paquete handler con el controlador de la entidad seleccionada.
// - Se debe generar la estructura request
// - Se debe generar la estructura del controlador que tenga como campo el servicio
// - Se debe generar la función que retorne el controlador
// - Se deben generar todos los métodos correspondientes a los endpoints

import (
	transactions "backpack-bcgow6-franco-niz/go-web/practica3/TM/internal/transactions"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	ID       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency"`
	Price    float64 `json:"price" binding:"required"`
	Emitter  string  `json:"emitter"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

type Transaction struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transaction {
	return &Transaction{
		service: s,
	}
}

func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		t, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, t)
	}
}

func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var message string
			if req.Code == "" {
				message = message + fmt.Sprintf("El campo Code es requerido. ")
			}

			if req.Price == 0 {
				message = message + fmt.Sprintf("El campo Price es requerido. ")
			}

			ctx.String(400, message)
			return
		}

		var message string
		if req.Currency == "" {
			message = message + fmt.Sprintf("El campo Currency es requerido. ")
		}
		if req.Emitter == "" {
			message = message + fmt.Sprintf("El campo Emitter es requerido. ")
		}
		if req.Receiver == "" {
			message = message + fmt.Sprintf("El campo Receiver es requerido. ")
		}
		if req.Date == "" {
			message = message + fmt.Sprintf("El campo Date es requerido. ")
		}
		if message != "" {
			ctx.String(400, message)
			return
		}

		t, err := c.service.Store(req.Code, req.Currency, req.Price, req.Emitter, req.Receiver, req.Date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, t)
	}
}

func (c *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "ID invalido",
			})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var message string
			if req.Code == "" {
				message = message + fmt.Sprintf("El campo Code es requerido. ")
			}

			if req.Price == 0 {
				message = message + fmt.Sprintf("El campo Price es requerido. ")
			}

			ctx.String(400, message)
			return
		}

		var message string
		if req.Currency == "" {
			message = message + fmt.Sprintf("El campo Currency es requerido. ")
		}
		if req.Emitter == "" {
			message = message + fmt.Sprintf("El campo Emitter es requerido. ")
		}
		if req.Receiver == "" {
			message = message + fmt.Sprintf("El campo Receiver es requerido. ")
		}
		if req.Date == "" {
			message = message + fmt.Sprintf("El campo Date es requerido. ")
		}
		if message != "" {
			ctx.String(400, message)
			return
		}

		t, err := c.service.Update(int(id), req.Code, req.Currency, req.Price, req.Emitter, req.Receiver, req.Date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, t)
	}
}

func (c *Transaction) PartialUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		if token != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "ID invalido",
			})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var message string
			if req.Code == "" {
				message = message + fmt.Sprintf("El campo Code es requerido. ")
			}
			if req.Price == 0 {
				message = message + fmt.Sprintf("El campo Price es requerido. ")
			}
			ctx.JSON(400, gin.H{
				"error": message,
			})
			return
		}

		t, err := c.service.PartialUpdate(int(id), req.Code, req.Price)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, t)

	}
}

func (c *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "token" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token invalido",
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "ID invalido",
			})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("La transaccion con id %d fue eliminada", id),
		})
	}
}

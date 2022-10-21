package handler

import (
	transactions "backpack-bcgow6-franco-niz/testing/practica3/TM/ej1/internal/transactions"
	"backpack-bcgow6-franco-niz/testing/practica3/TM/ej1/pkg/web"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	TokenError = "token invalido"
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

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (c *Transaction) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := tokenValidator(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return

		}

		t, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		if len(t) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, web.NotElements))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := tokenValidator(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			var message string
			if req.Code == "" {
				message = message + "El campo Code es requerido. "
			}

			if req.Price == 0 {
				message = message + "El campo Price es requerido. "
			}

			ctx.JSON(400, web.NewResponse(400, nil, message))
			return
		}

		var message string
		if req.Currency == "" {
			message = message + "El campo Currency es requerido. "
		}
		if req.Emitter == "" {
			message = message + "El campo Emitter es requerido. "
		}
		if req.Receiver == "" {
			message = message + "El campo Receiver es requerido. "
		}
		if req.Date == "" {
			message = message + "El campo Date es requerido. "
		}
		if message != "" {
			ctx.JSON(400, web.NewResponse(400, nil, message))
			return
		}

		t, err := c.service.Store(req.Code, req.Currency, req.Price, req.Emitter, req.Receiver, req.Date)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, t, ""))
	}
}

// UpdateTransactions godoc
// @Summary Update transactions
// @Tags Transactions
// @Description update transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to update"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := tokenValidator(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
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
				message = message + "El campo Code es requerido. "
			}

			if req.Price == 0 {
				message = message + "El campo Price es requerido. "
			}

			ctx.String(400, message)
			return
		}

		var message string
		if req.Currency == "" {
			message = message + "El campo Currency es requerido. "
		}
		if req.Emitter == "" {
			message = message + "El campo Emitter es requerido. "
		}
		if req.Receiver == "" {
			message = message + "El campo Receiver es requerido. "
		}
		if req.Date == "" {
			message = message + "El campo Date es requerido. "
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

// PartialUpdateTransactions godoc
// @Summary PartialUpdate transactions
// @Tags Transactions
// @Description partial update transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to update partiacially"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *Transaction) PartialUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := tokenValidator(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
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
				message = message + "El campo Code es requerido. "
			}
			if req.Price == 0 {
				message = message + "El campo Price es requerido. "
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

// DeleteTransactions godoc
// @Summary Delete transactions
// @Tags Transactions
// @Description delete transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to delete"
// @Success 200 {object} web.Response
// @Router /transactions [post]
func (c *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if err := tokenValidator(ctx); err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, web.InvalidToken))
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El ID es invalido"))
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, fmt.Sprintf("La transaccion con id %d fue eliminada", id), ""))
	}
}

func tokenValidator(ctx *gin.Context) error {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {

		return fmt.Errorf(TokenError)

	}
	return nil
}

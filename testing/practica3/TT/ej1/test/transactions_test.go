package test

import (
	"backpack-bcgow6-franco-niz/testing/practica3/TT/ej1/cmd/server/handler"
	transactions "backpack-bcgow6-franco-niz/testing/practica3/TT/ej1/internal/transactions"
	"backpack-bcgow6-franco-niz/testing/practica3/TT/ej1/pkg/store"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	db := store.New(store.FileType, "../transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.POST("/:id", t.Update())
	tr.DELETE("/:id", t.Delete())
	return r

}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "my_token")

	return req, httptest.NewRecorder()
}

func Test_GetTransaction_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
	//err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	//assert.Nil(t, err)
	//assert.True(t, len(objRes.Data) > 0)
}

func Test_SaveTransaction_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
		"code": "algunCode","currency": "dolares","price": 405.2,"emitter": "MeLi","receiver": "MeLi","date": "09/10"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_UpdateTransaction_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/transactions/3", `{
		"code": "algunCodeNew","currency": "dolares","price": 405.2,"emitter": "MeLi","receiver": "MeLi","date": "09/10"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_DeleteTransaction_OK(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/transactions/4", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

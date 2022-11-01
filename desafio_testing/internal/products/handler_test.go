package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ProductsRoute = "/api/v1/products"
)

func createServer() *gin.Engine {
	r := gin.Default()

	//router.MapRoutes(r)
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	r.GET(ProductsRoute+"/", handler.GetProducts)

	return r
}

func createServerRepositoryError() *gin.Engine {
	r := gin.Default()

	//router.MapRoutes(r)
	repo := NewRepositoryError()
	service := NewService(&repo)
	handler := NewHandler(service)

	r.GET(ProductsRoute+"/", handler.GetProducts)

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func Test_GetProduct(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, ProductsRoute+"/", "")
	params := req.URL.Query()
	params.Add("seller_id", "FEX112AC")
	req.URL.RawQuery = params.Encode()
	//req.Header.Add("seller_id", "FEX112AC")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func Test_GetProduct_error_queryparam(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodGet, ProductsRoute+"/", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 400, rr.Code)
}

func Test_GetProduct_error_repository(t *testing.T) {
	r := createServerRepositoryError()

	req, rr := createRequestTest(http.MethodGet, ProductsRoute+"/", "")
	params := req.URL.Query()
	params.Add("seller_id", "FEX112AC")
	req.URL.RawQuery = params.Encode()

	r.ServeHTTP(rr, req)

	assert.Equal(t, 500, rr.Code)
}

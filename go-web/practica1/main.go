package main

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

// Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
// Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
// Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
// Genera un handler para el endpoint llamado “GetAll”.
// Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.

// Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los
// campos.
// Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
// Luego genera la lógica de filtrado de nuestro array.
// Devolver por el endpoint el array filtrado.

// Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática. Utilizando path parameters el endpoint
// debería ser /temática/:id (recuerda que siempre tiene que ser en plural la temática). Una vez recibido el id devuelve la posición
// correspondiente.
// Genera una nueva ruta.
// Genera un handler para la ruta creada.
// Dentro del handler busca el item que necesitas.
// Devuelve el item según el id.
// Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var transactions []Transactions
var transactionsMap map[string]string

func JSONToSlice() {
	content, err := os.ReadFile("./transactions.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &transactions)
	if err != nil {
		panic(err)
	}
}

func GetAll(c *gin.Context) {

	idFiltered, _ := strconv.Atoi(c.Query("id"))
	priceFiltered, _ := strconv.ParseFloat(c.Query("price"), 64)
	filteredTransaction := Transactions{
		ID:       idFiltered,
		Code:     c.Query("code"),
		Currency: c.Query("currency"),
		Price:    priceFiltered,
		Emitter:  c.Query("emitter"),
		Receiver: c.Query("receiver"),
		Date:     c.Query("date"),
	}

	transactions = FilteredSlice(filteredTransaction)

	c.JSON(200, gin.H{
		"message":      "Hola Franco!",
		"transactions": transactions,
	})

}

func FilteredSlice(t Transactions) (trs []Transactions) {

	if t.ID != 0 && t.Code != "" {
		for _, v := range transactions {
			if v.Code == t.Code && t.ID == v.ID {
				trs = append(trs, v)
			}
		}
	} else if t.ID != 0 {
		for _, v := range transactions {
			if t.ID == v.ID {
				trs = append(trs, v)
			}
		}
	} else if t.Code != "" {
		for _, v := range transactions {
			if v.Code == t.Code {
				trs = append(trs, v)
			}
		}
	} else {
		trs = transactions
	}

	return
}

func FileteredID(c *gin.Context) {
	trID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	var t Transactions
	for _, v := range transactions {
		fmt.Printf("v: %v\n", v)
		if v.ID == trID {
			fmt.Printf("t: %v\n", t)
			t = v
			break
		}
	}

	fmt.Printf("t: %v\n", t)
	if t.ID == 0 {
		c.String(http.StatusNotFound, "El id %d no se encuentra registrado", trID)
	} else {
		c.String(200, "El id %d es: %v", trID, t)
	}
}

type Transactions struct {
	ID       int
	Code     string
	Currency string
	Price    float64
	Emitter  string
	Receiver string
	Date     string
}

func main() {

	//GetAll()
	JSONToSlice()

	router := gin.Default()
	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", FileteredID)
	router.Run()
}

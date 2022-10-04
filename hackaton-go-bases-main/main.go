package main

import (
	"fmt"
	"io"
	"os"

	"hackaton-go-bases-main/internal/service"
)

func HandleError() {
	err := recover()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	file, err := os.OpenFile("./tickets.csv", os.O_RDWR, 0644)
	defer fmt.Println("Ejecucion terminada")
	defer file.Close()
	defer HandleError()
	if err != nil {
		panic(err)
	}

	booking, err := service.NewBookings(tickets, file)
	if err != nil && err != io.EOF {
		panic(err)
	}

	newTicket := service.Ticket{
		Id:          1030,
		Names:       "Franco Niz",
		Email:       "franco.niz@mercadolibre.com",
		Destination: "Brasil",
		Date:        "9:40",
		Price:       40000,
	}

	t, err := booking.Create(newTicket)
	if err != nil {
		panic(err)
	}

	t, _, err = booking.Read(1030)
	if err != nil {
		panic(err)
	}
	fmt.Printf("t: %v\n", t)

	t.Names = "Franco Damian Niz"
	err = booking.Update(t.Id, t)
	if err != nil {
		panic(err)
	}

	t, _, err = booking.Read(1030)
	if err != nil {
		panic(err)
	}
	fmt.Printf("t: %v\n", t)

	booking.Delete(1030)
	t, _, err = booking.Read(1030)
	if err != nil {
		panic(err)
	}
	fmt.Printf("t: %v\n", t)
}

package service

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, int, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) error
	// Delete delete a Ticket by id
	Delete(id int) error
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket, file *os.File) (Bookings, error) {
	_, seekErr := file.Seek(0, 0)
	if seekErr != nil {
		return &bookings{}, seekErr
	}
	csvReader := csv.NewReader(file)

	var t []string
	var err error
	t, err = csvReader.Read()

	for err != io.EOF { // Lee los tickets hasta llegar llegar al final

		//fmt.Printf("t: %v\n", t)
		id, _ := strconv.Atoi(t[0])
		price, _ := strconv.Atoi(t[5])

		ticket := Ticket{
			Id:          id, // v es equivalente a Atoi(t[0])
			Names:       t[1],
			Email:       t[2],
			Destination: t[3],
			Date:        t[4],
			Price:       price,
		}
		//fmt.Printf("ticket: %v\n", ticket)
		//fmt.Printf("ticket: %v\n", ticket)
		Tickets = append(Tickets, ticket)

		t, err = csvReader.Read()
	}

	//fmt.Println("llegue")
	return &bookings{Tickets: Tickets}, err
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	newId := t.Id

	i := 0
	for i < len(b.Tickets) && newId != b.Tickets[i].Id {
		i++
	}
	if i < len(b.Tickets) {
		if newId == b.Tickets[i].Id {
			return Ticket{}, errors.New("The ID exists in the data base.")
		}
	}

	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, int, error) {

	for i := 0; i < len(b.Tickets); i++ { // En realidad, al tener tickets con id ordenado, solo
		actualTicket := b.Tickets[i] // habria que fijarse que el id < len(b.Tickets) para saber si este existe
		if actualTicket.Id == id {   // pero aca estoy suponiendo que los tickets no estan necesariamente en orden de id
			return actualTicket, i, nil
		}
	}

	return Ticket{}, 0, errors.New("The id does not correspond to a ticket")
}

func (b *bookings) Update(id int, t Ticket) error {

	_, pos, err := b.Read(id)
	if err != nil {
		return err
	}

	b.Tickets[pos].Id = t.Id
	b.Tickets[pos].Email = t.Email
	b.Tickets[pos].Destination = t.Destination
	b.Tickets[pos].Names = t.Names
	b.Tickets[pos].Date = t.Date
	b.Tickets[pos].Price = t.Price
	return nil
}

func (b *bookings) Delete(id int) error {

	_, pos, err := b.Read(id)
	if err != nil {
		return err
	}

	b.Tickets[pos] = b.Tickets[len(b.Tickets)-1]
	b.Tickets = b.Tickets[:len(b.Tickets)-1]

	return nil
}

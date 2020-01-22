package api

import (
	"fmt"

	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

func NewTicketRepository() TicketRepository {
	return TicketRepository{}
}

type TicketRepository struct {}

func (r TicketRepository) Read(ticketID int) (domain.Ticket, error) {
	request := `
		curl -X GET https://tickets.api/{id}
	`
	fmt.Println(request)

	return domain.Ticket{ticketID}, nil
}
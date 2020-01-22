package database

import (
	"fmt"

	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

func NewTicketRepository() TicketRepository {
	return TicketRepository{}
}

type TicketRepository struct {}

func (r TicketRepository) Read(ticketID int) (domain.Ticket, error) {
	sql := `
		SELECT * FROM tickets WHERE id = ?
	`
	fmt.Println(sql)

	return domain.Ticket{ticketID}, nil
}

func (r TicketRepository) Write(ticketID int) error {
	sql := `
		INSERT INTO tickets(id) VALUES (?)
	`
	fmt.Println(sql)

	return nil
}
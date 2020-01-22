package main

import (
	"fmt"

	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

type TicketGetter interface {
	GetTicketByID(ticketID int) (domain.Ticket, error)
}

type TicketSaver interface {
	SaveTicketWithID(ticketID int) error
}

func NewApp(ticketGetter TicketGetter, ticketSaver TicketSaver) App {
	return App{
		ticketGetter: ticketGetter,
		ticketSaver: ticketSaver,
	}
}

type App struct{
	ticketGetter TicketGetter
	ticketSaver TicketSaver
}

func (a App) Handle(action string, ticketID int) {
	switch action {
	case "get":
		ticket, err := a.ticketGetter.GetTicketByID(ticketID)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ticket)
		return
	case "save":
		err := a.ticketSaver.SaveTicketWithID(ticketID)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}

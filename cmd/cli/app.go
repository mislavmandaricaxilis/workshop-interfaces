package main

import (
	"fmt"

	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

type TicketGetter interface {
	GetTicketByID(ticketID int) (domain.Ticket, error)
}

func NewApp(ticketGetter TicketGetter) App {
	return App{
		ticketGetter: ticketGetter,
	}
}

type App struct{
	ticketGetter TicketGetter
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
		panic("Not implemented")
	}
}

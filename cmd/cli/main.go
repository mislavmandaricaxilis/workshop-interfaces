package main

import (
	"flag"
	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

var action = flag.String("action", "get", "action to do on a ticket")
var ticketID = flag.Int("ticket", 0, "unique identifier for ticket")

func main() {
	flag.Parse()

	app := bootstrap()
	app.Handle(*action, *ticketID)
}

func bootstrap() App {
	ticketService := domain.NewTicketService()
	return NewApp(ticketService)
}
package main

import (
	"flag"

	"github.com/mislavmandaricaxilis/workshop-interfaces-dependency"
	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/database"
	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/services"
)

var action = flag.String("action", "get", "action to do on a ticket")
var ticketID = flag.Int("ticket", 0, "unique identifier for ticket")

func main() {
	flag.Parse()

	app := bootstrap()
	app.Handle(*action, *ticketID)
}

func bootstrap() App {
	logger := workshop_interfaces_dependency.NewLoggerOneWay()
	ticketRepository := database.NewTicketRepository()
	ticketService := services.NewTicketService(ticketRepository, logger)
	return NewApp(ticketService)
}
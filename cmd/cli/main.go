package main

import (
	"flag"
	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"

	"github.com/mislavmandaricaxilis/workshop-interfaces-dependency/v2"
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
	loggerV2 := workshop_interfaces_dependency.NewLoggerOneWay()
	logger := domain.NewLogger(loggerV2)
	ticketRepository := database.NewTicketRepository()
	ticketService := services.NewTicketService(ticketRepository, ticketRepository, logger)
	return NewApp(ticketService, ticketService)
}
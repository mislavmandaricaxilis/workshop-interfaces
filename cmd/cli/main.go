package main

import (
	"flag"
)

var action = flag.String("action", "get", "action to do on a ticket")
var ticketID = flag.Int("ticket", 0, "unique identifier for ticket")

func main() {
	flag.Parse()

	app := NewApp()

	app.Handle(*action, *ticketID)
}
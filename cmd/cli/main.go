package main

import (
	"flag"
)

var ticketID = flag.Int("ticket", 0, "unique identifier for ticket")

func main() {
	flag.Parse()

	app := NewApp()

	app.Handle(*ticketID)
}
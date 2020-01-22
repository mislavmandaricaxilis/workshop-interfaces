package main

import (
	"flag"
	"fmt"
)

var ticketID = flag.Int("ticket", 0, "unique identifier for ticket")

func main() {
	flag.Parse()

	fmt.Println(*ticketID)
}
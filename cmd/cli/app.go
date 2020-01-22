package main

import "fmt"

func NewApp() App {
	return App{}
}

type App struct{}

func (a App) Handle(ticketID int) {
	fmt.Println(ticketID)
}

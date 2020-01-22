package main

import "fmt"

func NewApp() App {
	return App{}
}

type App struct{}

func (a App) Handle(action string, ticketID int) {
	fmt.Println(action)
	fmt.Println(ticketID)
}

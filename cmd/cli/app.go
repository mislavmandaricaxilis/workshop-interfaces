package main

func NewApp() App {
	return App{}
}

type App struct{
}

func (a App) Handle(action string, ticketID int) {
	switch action {
	case "get":
		panic("Not implemented")
	case "save":
		panic("Not implemented")
	}
}

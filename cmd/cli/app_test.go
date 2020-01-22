package main

import (
	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
	"testing"
)

func Test_WhenActionIsGet_Handle_ShouldNotPanic(t *testing.T) {
	// Assert
	defer func() {
		if r := recover(); r != nil {
			t.Fail()
		}
	}()

	// Arrange
	app := NewApp(mockTicketGetter{})

	// Act
	app.Handle("get", 1)
}

type mockTicketGetter struct {}

func (g mockTicketGetter) GetTicketByID(ticketID int) (domain.Ticket, error) {
	return domain.Ticket{ticketID}, nil
}

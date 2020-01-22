package services

import (
	"errors"

	"github.com/mislavmandaricaxilis/workshop-interfaces/internal/domain"
)

type Logger interface {
	Log(s string)
}

type DatabaseReader interface {
	Read(ticketID int) (domain.Ticket, error)
}

func NewTicketService(databaseReader DatabaseReader, logger Logger) TicketService {
	return TicketService{
		databaseReader: databaseReader,
		logger: logger,
	}
}

type TicketService struct {
	databaseReader DatabaseReader
	logger Logger
}

func (s TicketService) GetTicketByID(ticketID int) (domain.Ticket, error) {
	if ticketID == 0 {
		s.logger.Log("ticket ID was not provided when getting ticket")
		return domain.Ticket{}, nil
	}

	ticket, err := s.databaseReader.Read(ticketID)
	if err != nil {
		return domain.Ticket{}, errors.New("error with reading from database")
	}

	return ticket, nil
}
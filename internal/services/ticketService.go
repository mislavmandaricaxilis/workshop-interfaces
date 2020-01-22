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

type DatabaseWriter interface {
	Write(ticketID int) error
}

func NewTicketService(databaseReader DatabaseReader, databaseWriter DatabaseWriter, logger Logger) TicketService {
	return TicketService{
		databaseReader: databaseReader,
		databaseWriter: databaseWriter,
		logger: logger,
	}
}

type TicketService struct {
	databaseReader DatabaseReader
	databaseWriter DatabaseWriter
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

func (s TicketService) SaveTicketWithID(ticketID int) error {
	if ticketID == 0 {
		return  errors.New("can't save ticket with ID 0")
	}

	err := s.databaseWriter.Write(ticketID)
	if err != nil {
		return errors.New("error with writing to database")
	}

	return nil
}
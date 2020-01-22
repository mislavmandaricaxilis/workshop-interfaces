package domain

func NewTicketService() TicketService {
	return TicketService{}
}

type TicketService struct {}

func (s TicketService) GetTicketByID(ticketID int) (Ticket, error) {
	return Ticket{Id:ticketID}, nil
}
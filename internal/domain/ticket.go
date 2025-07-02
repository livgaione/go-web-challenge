package domain

type Ticket struct {
	Id         int              `json:"id"`
	Attributes TicketAttributes `json:"attributes"`
}

type TicketAttributes struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Hour    string  `json:"hour"`
	Price   float64 `json:"price"`
}

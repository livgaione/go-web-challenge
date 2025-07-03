package repository

import (
	"context"
	"log"
	"tickets/internal/domain"
)

type RepositoryTicket interface {
	Get(ctx context.Context) (t map[int]domain.TicketAttributes, err error)
	GetTicketByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error)
	GetAverage(country string) (int, error)
}

type RepositoryTicketMap struct {
	db     map[int]domain.TicketAttributes
	lastId int
}

func NewRepositoryTicket(dbFile map[int]domain.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     dbFile,
		lastId: lastId,
	}
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]domain.TicketAttributes, err error) {
	t = make(map[int]domain.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return t, nil
}

func (r *RepositoryTicketMap) GetTicketByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	t = make(map[int]domain.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	log.Println(t)
	return t, nil
}

func (r *RepositoryTicketMap) GetAverage(country string) (int, error) {
	var count int
	var avg int

	for _, k := range r.db {
		if k.Country == country {
			count++
		}

		avg = (count / len(k.Country))

	}

	return avg, nil

}

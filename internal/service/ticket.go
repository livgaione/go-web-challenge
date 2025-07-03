package service

import (
	"context"
	"tickets/internal/domain"
	"tickets/internal/repository"
	"tickets/pkg/apperrors"
)

type ServiceTicket interface {
	GetAll(ctx context.Context) (t map[int]domain.TicketAttributes, err error)
	GetTicketByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error)
	GetAverage(country string) (int, error)
}

type ServiceTicketDefault struct {
	rp repository.RepositoryTicket
}

func NewServiceTicketDefault(rp repository.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

func (s *ServiceTicketDefault) GetAll(ctx context.Context) (t map[int]domain.TicketAttributes, err error) {
	tickets, err := s.rp.Get(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *ServiceTicketDefault) GetTicketByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	tickets, err := s.rp.GetTicketByDestinationCountry(country)
	if err != nil {
		return nil, apperrors.ErrResourceNotExists
	}
	return tickets, nil
}

func (s *ServiceTicketDefault) GetAverage(country string) (int, error) {
	tickets, err := s.rp.GetAverage(country)
	if err != nil {
		return 0, apperrors.ErrResourceNotExists
	}
	return tickets, nil
}

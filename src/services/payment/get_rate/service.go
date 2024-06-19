package get_rate

import (
	"context"
)

type service struct {
	rateGateway RateGateway
}

type Service interface {
	Execute(context.Context) (Rate, error)
}

func NewService(rateGateway RateGateway) Service {
	return &service{
		rateGateway: rateGateway,
	}
}

func (s service) Execute(ctx context.Context) (Rate, error) {
	rate, err := s.rateGateway.GetRate(ctx)
	if err != nil {
		return rate, err
	}
	return rate, nil
}

package get_rate

import "context"

type RateGateway interface {
	GetRate(ctx context.Context) (Rate, error)
}

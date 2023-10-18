package services

import (
	"github.com/ricardokovalski/treinando-go-lang/src/entities"
)

func FetchPlan() entities.Plan {
	plan := entities.Plan{
		Identifier: "plan_start_2023",
		Price:      289.79,
		Range:      4,
		Limits: []entities.Limit{
			{
				Key:   "dfes",
				Value: 200,
			},
			{
				Key:   "coupons",
				Value: 2000,
			},
		},
	}

	return plan
}

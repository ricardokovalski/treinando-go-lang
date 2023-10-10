package services

import "github.com/ricardokovalski/treinando-go-lang/src/entities"

func FetchPlans() []entities.Plan {

	collection := []entities.Plan{
		{
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
		},
		{
			Identifier: "plan_basic_2023",
			Price:      299.79,
			Range:      5,
			Limits: []entities.Limit{
				{
					Key:   "dfes",
					Value: 300,
				},
				{
					Key:   "coupons",
					Value: 3000,
				},
			},
		},
	}
	return collection
}

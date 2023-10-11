package services

import (
	"github.com/ricardokovalski/treinando-go-lang/src/entities"
)

func AppendPlan(plans []entities.Plan, newPlan entities.Plan) []entities.Plan {
	plans = append(plans, newPlan)
	return plans
}

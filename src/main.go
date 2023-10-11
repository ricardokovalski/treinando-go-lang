package main

import (
	"fmt"

	"github.com/ricardokovalski/treinando-go-lang/src/entities"
	"github.com/ricardokovalski/treinando-go-lang/src/services"
)

func main() {
	// obtendo um plano mockado
	plan := services.FetchPlan()
	// encodando o plano retornando
	json, _ := services.EncodePlan(plan)
	// decodificando o json em uma estruct de Plan
	planDecode, _ := services.DecodePlan(json)

	fmt.Println(plan)
	fmt.Println(json)
	fmt.Println(planDecode)

	// obtendo vários planos
	plans := services.FetchPlans()
	fmt.Printf("%+v\n", plans)

	newPlan1 := entities.Plan{
		Identifier: "plan_analytical_2023",
		Price:      400.98,
		Range:      7,
		Limits: []entities.Limit{
			{
				Key:   "dfes",
				Value: 400,
			},
			{
				Key:   "coupons",
				Value: 4000,
			},
		},
	}

	plans = services.AppendPlan(plans, newPlan1)
	fmt.Printf("%+v\n", plans)

}

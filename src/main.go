package main

import (
	"context"

	"github.com/ricardokovalski/treinando-go-lang/src/adapters/payment/rate"
	"github.com/ricardokovalski/treinando-go-lang/src/services/payment/get_rate"
	"github.com/rs/zerolog/log"
)

func main() {

	ctx := context.Background()
	ctx = log.Logger.WithContext(ctx)

	service := get_rate.NewService(
		rate.GetClient(
			"http://localhost:8080",
		),
	)

	service.Execute(ctx)

	// obtendo um plano mockado
	//plan := services.FetchPlan()
	// encodando o plano retornando
	//json, _ := utils.Encode(plan)
	// decodificando o json em uma estruct de Plan
	//planDecode, _ := services.DecodePlan(json)

	// fmt.Println(plan)
	// fmt.Println(json)
	// fmt.Println(planDecode)

	// obtendo vários planos
	// fmt.Println("Obtendo todos os planos")
	// plans := services.FetchPlans()
	// fmt.Printf("%+v\n", plans)

	// newPlan1 := entities.Plan{
	// 	Identifier: "plan_analytical_2023",
	// 	Price:      400.98,
	// 	Range:      7,
	// 	Limits: []entities.Limit{
	// 		{
	// 			Key:   "dfes",
	// 			Value: 400,
	// 		},
	// 		{
	// 			Key:   "coupons",
	// 			Value: 4000,
	// 		},
	// 	},
	// }

	// fmt.Println("Add novo plano")
	// plans = services.AppendPlan(plans, newPlan1)

	// fmt.Println("Encodar a lista de planos")
	// jsonCollection, _ := utils.Encode(plans)
	// fmt.Println(jsonCollection)

	// fmt.Println("Decodando a lista de planos")
	// var jsonNew = []entities.Plan{}
	// utils.Decode(&jsonNew, jsonCollection)
	//fmt.Printf("%+v\n", test)
	//fmt.Println(jsonNew)
}

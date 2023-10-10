package main

import (
	"fmt"

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
}

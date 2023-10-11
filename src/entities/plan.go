package entities

type Plan struct {
	Identifier string  `json:"identifier"`
	Price      float64 `json:"price"`
	Range      int64   `json:"range"`
	Limits     []Limit `json:"limits"`
}

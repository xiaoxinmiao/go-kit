package base

import (
	"math"
)

const (
	DefaultRoundDigit    = 2
	DefaultRoundStrategy = "round"
	DefaultCurrency      = "CYN"
)

var (
	DefaultPriceSetting = &PriceSetting{
		RoundDigit:    DefaultRoundDigit,
		RoundStrategy: DefaultRoundStrategy,
		Currency:      DefaultCurrency,
	}
)

type PriceSetting struct {
	RoundDigit    int    `json:"roundDigit"`
	RoundStrategy string `json:"roundStrategy"`
	Currency      string `json:"currency"`
}

func ToFixed(num float64, priceSetting *PriceSetting) float64 {
	switch priceSetting.RoundStrategy {
	case "ceil", "Ceil":
		output := math.Pow(10, float64(priceSetting.RoundDigit))
		return math.Ceil(num*output) / output
	case "floor", "Floor":
		output := math.Pow(10, float64(priceSetting.RoundDigit))
		return math.Floor(num*output) / output
	default:
		output := math.Pow(10, float64(priceSetting.RoundDigit))
		return float64(Round(num*output)) / output
	}
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

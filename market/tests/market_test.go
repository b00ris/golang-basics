package tests

import (
	"lessons/basics/market"
	"testing"
)

func TestCalcOrder(t *testing.T) {
	Initialize()

	var order = market.Order{
		User: "Gates", Products: []string{"Whisky"},
	}

	var price = market.CalcOrder(order)
	t.Log(price)
	t.Log(market.Users["Gates"])
}

func Initialize() {
	market.Users["Kevin"] = 500
	market.Users["Gates"] = 1_000_000_000
	market.Users["Ford"] = 500_000_000

	market.Products["Pineapple"] = 15.0
	market.Products["Discount"] = 50.0
	market.Products["Whisky"] = 5_000.0
	market.Products["Meat"] = 125.0
	market.Products["Chocolate"] = 35.5
}

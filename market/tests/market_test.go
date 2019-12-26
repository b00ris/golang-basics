package tests

import (
	"lessons/basics/market"
	"math/rand"
	"testing"
)

const (
	ASC = iota
	DESC
	MANY
)

func TestCalcOrder(t *testing.T) {
	products := InitProducts()
	users := InitUsers()
	orders := InitOrder(products, users)

	market.CalcOrder(orders[1], users)

	market.PrintUsers(DESC, users)

	//t.Log(price)
	//t.Log(users["Gates"])
}

func BenchmarkCalcOrder(b *testing.B) {
	products := InitProducts()
	users := InitUsers()
	orders := InitOrder(products, users)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		market.CalcOrder(orders[rand.Intn(100)], users)
	}
}

func BenchmarkCalcOrderSpeedCache(b *testing.B) {
	products := InitProducts()
	users := InitUsers()
	orders := InitOrder(products, users)

	cache := map[string]float32{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		market.CalcOrderWithSpeedCache(orders[rand.Intn(100)], users, cache)
	}
}

func InitUsers() map[string]float32 {
	return map[string]float32{
		"Kevin": 1_500_000_000,
		"Gates": 1_000_000_000,
		"Ford":  500_000_000,
	}
}

func InitProducts() []market.Product {
	return []market.Product{
		{"Pineapple", 15.0},
		{"Discount", 50.0},
		{"Meat", 125.0},
		{"Whisky", 5_000.0},
		{"Pineapple", 15.0},
		{"Chocolate", 35.5},
		{"Banana", 25.25},
		{"Mango", 10.0},
		{"Marakua", 123.0},
		{"Potato", 2.0},
		{"Tea", 10.0},
		{"Coffee", 100.0},
	}
}
func InitOrder(products []market.Product, users map[string]float32) []market.Order {
	keyUsers := make([]string, 0, len(users))
	for k := range users {
		keyUsers = append(keyUsers, k)
	}

	orders := make([]market.Order, 100)
	for i := 0; i < 100; i++ {
		orders[i] = getOrder(keyUsers, products)
	}

	return orders
}

func getOrder(userList []string, productList []market.Product) market.Order {
	length := rand.Intn(100)
	products := make([]market.Product, length)

	for i := 0; i < length; i++ {
		products[i] = productList[rand.Intn(len(productList))]
	}

	return market.Order{
		User:     userList[rand.Intn(len(userList))],
		Products: products,
	}
}

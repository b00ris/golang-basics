package main

import (
	"fmt"
	"sort"
)

const (
	ASC  = 1
	DESC = 2
	MANY = 3
)

var products = map[string]float32{
	"Bread":  9.99,
	"Papaya": 25.45,
	"Cat":    1.0,
}

var users = map[string]float32{
	"Вася":  300.0,
	"Петя":  30000000.0,
	"Яоурт": 500.0,
}

func main() {
	addProduct("Dog", 199.99)
	userOrder := []string{"Bread", "Cat"}
	price := calcOrder("Вася", userOrder)
	fmt.Println(price)
	//fmt.Println(users)
	printUsers(DESC)
}

func calcOrder(user string, order []string) float32 {
	var result float32
	for _, v := range order {
		price, okay := products[v]
		if okay != false {
			result += price
		}
	}

	if users[user] >= result {
		users[user] -= result
		return result
	}
	return 0
}

func printUsers(typeSort int) {
	switch typeSort {
	case ASC:
		keys := make([]string, 0, len(users))
		for k := range users {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Println(k, products[k])
		}
	case DESC:
		keys := make([]string, 0, len(users))
		for k := range users {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		for _, k := range keys {
			fmt.Println(k, products[k])
		}
	case MANY:

	}
}

func addProduct(product string, price float32) {
	_, okay := products[product]
	if okay == false {
		products[product] = price
	}
}

func updateProduct(product string, price float32) {
	_, okay := products[product]
	if okay != false {
		products[product] = price
	}
}

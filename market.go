package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	ASC  = 1
	DESC = 2
	MANY = 3
)

type pair struct {
	key   string
	value float32
}

var cacheOrder = map[string]float32{}

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

func main3() {
	addProduct("Dog", 199.99)
	userOrder := []string{"Bread", "Cat"}
	price := calcOrder("Вася", userOrder)
	fmt.Println(price)
	userOrder2 := []string{"Cat", "Bread"}
	price = calcOrder("Вася", userOrder2)
	fmt.Println(price)
	//fmt.Println(users)
	printUsers(MANY)
}

func calcOrder(user string, order []string) float32 {
	var result float32
	sort.Strings(order)
	orderKey := strings.Join(order, " ")
	elem, okay := cacheOrder[orderKey]
	if okay != false {
		fmt.Println("USE CACHE")
		result = elem
	} else {
		for _, v := range order {
			price, okay := products[v]
			if okay != false {
				result += price
			}
		}
		cacheOrder[orderKey] = result
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
			fmt.Println(k, users[k])
		}
	case DESC:
		keys := make([]string, 0, len(users))
		for k := range users {
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		for _, k := range keys {
			fmt.Println(k, users[k])
		}
	case MANY:
		values := make([]pair, len(users))
		i := 0
		for k, v := range users {
			values[i] = pair{k, v}
			i++
		}
		sort.Slice(values, func(i, j int) bool {
			return values[i].value > values[j].value
		})

		for _, v := range values {
			fmt.Println(v.key, v.value)
		}
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

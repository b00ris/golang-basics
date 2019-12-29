package market

import (
	"fmt"
	"sort"
)

type sortDirection int

const (
	ASC sortDirection = iota
	DESC
	MANY
)

type pair struct {
	key   string
	value float32
}

type Product struct {
	Name  string
	Price float32
}
type Order struct {
	User     string
	Products []Product
}

func CalcOrder(order Order, users map[string]float32) float32 {
	var total float32
	for _, product := range order.Products {
		total += product.Price
	}
	value, okay := users[order.User]
	if okay == true && value >= total {
		users[order.User] -= total
	}
	return total
}

func CalcOrderWithSpeedCache(order Order, users map[string]float32, cache map[string]float32) float32 {
	var total float32

	// jekamas: кэширование лушче вынести в отдельную функцию
	orderKey := ""
	for _, v := range order.Products {
		orderKey += v.Name
	}

	var ok bool
	if total, ok = cache[orderKey]; !ok {
		for _, product := range order.Products {
			total += product.Price
		}
		cache[orderKey] = total
	}

	if value, ok := users[order.User]; ok && value >= total {
		users[order.User] -= total
	} else {
		// fmt.Println(order.User, " не имеет требуемой суммы")
	}

	return total
}

func PrintUsers(direction sortDirection, users map[string]float32) {
	switch direction {
	case ASC:
		keys := make([]string, 0, len(users))
		for k := range users {
			keys = append(keys, k)
		}
		sort.Strings(keys)
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

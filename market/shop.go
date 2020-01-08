package market

import (
	"fmt"
	"sort"
)

// Инициализация магазина
func ShopInit(products map[string]Product, users map[string]User) *Shop {
	return &Shop{Products: products,
		Users: users,
		Cache: map[string]float32{},
	}
}

// Подсчет стоимости заказа
func (shop *Shop) CalcOrder(order Order, useCache bool) float32 {
	var total float32

	user, okay := shop.Users[order.User]
	if okay {
		var cache CacheInfo

		if useCache {
			cache = getCacheTotal(order, shop, user.Status)
		}
		if cache.IsHas == true {
			total = cache.Total
		} else {
			var product Product

			for _, productName := range order.Products {
				product = shop.Products[productName]
				total += product.Price * calcCoefficientByStatus(user.Status, product.Status)
			}

			var kit Kit
			for _, kitName := range order.Kits {
				kit = shop.Kits[kitName]
				mainProduct := shop.Products[kit.MainProduct]
				additionalProduct := shop.Products[kit.AdditionalProduct]
				mainPrice := mainProduct.Price * calcCoefficientByStatus(user.Status, mainProduct.Status)
				additionalPrice := additionalProduct.Price * calcCoefficientByStatus(user.Status, additionalProduct.Status)

				total += (mainPrice + additionalPrice) * kit.Discount
			}
			if useCache {
				shop.Cache[cache.Key] = total
			}
		}
		if user.Bill >= total {
			user.Bill -= total
			shop.Users[order.User] = user
		}
	} else {
		fmt.Println("User not found")
	}
	return total
}

// Вывод состояния пользователей с различной сортировкой
func (shop *Shop) PrintUsers(direction sortDirection) {
	switch direction {
	case ASC:
		keys := make([]string, 0, len(shop.Users))
		for k := range shop.Users {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Println(k, shop.Users[k])
		}
	case MANY:
		values := make([]pair, len(shop.Users))
		i := 0
		for k, v := range shop.Users {
			values[i] = pair{k, v.Bill}
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

//// Подсчет стоимости заказа с использованием кэширования вычислений
//func (shop *Shop)CalcOrderWithSpeedCache(order Order) float32 {
//	var total float32
//
//	user, ok := shop.Users[order.User]
//	if ok {
//		//total = getTotal(order, shop, user.Status)
//		if user.Bill >= total {
//			user.Bill -= total
//			shop.Users[order.User] = user
//		} else {
//			// fmt.Println(order.User, " не имеет требуемой суммы")
//		}
//	}
//
//	return total
//}

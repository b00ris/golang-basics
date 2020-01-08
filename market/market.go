package market

// Подсчет коэффицента для особых пользователей и особых товаров
func calcCoefficientByStatus(userStatus typeStatus, productStatus typeStatus) float32 {
	var coefficient float32 = 1.0
	switch userStatus {
	case NORMAL:
		switch productStatus {
		case PREMIUM:
			coefficient = 1.5
		}
	case PREMIUM:
		switch productStatus {
		case NORMAL:
			coefficient = 0.95
		case PREMIUM:
			coefficient = 0.80
		}
	}
	return coefficient
}

// Получение кэшированного значения вычислений
func getCacheTotal(order Order, shop *Shop, userStatus typeStatus) CacheInfo {
	orderKey := string(userStatus)
	for _, v := range order.Products {
		orderKey += v
	}

	for _, v := range order.Kits {
		orderKey += v
	}
	total, ok := shop.Cache[orderKey]

	return CacheInfo{
		IsHas: ok,
		Total: total,
		Key:   orderKey,
	}
}

//  Создание набора
func CreateKit(main string, additional string, probe Probe, discount float32) Kit {
	return Kit{
		MainProduct:       main,
		AdditionalProduct: additional,
		Probe:             probe,
		Discount:          1 - (discount / 100),
	}
}

//  Обновление скидки на набор
func (shop *Shop) UpdateKit(name string, discount float32) {
	kit := shop.Kits[name]
	kit.Discount = 1 - (discount / 100)
	shop.Kits[name] = kit
}

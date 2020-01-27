package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	"sort"
)

type MyOrder shop_competition.Order

//Подсчет суммы заказа без учета скидок и наценок
func (shop Shop) CalculateOrder(order shop_competition.Order) (float32, error) {
	if order.Products == nil &&
		order.Bundles == nil {
		return 0, errors.New("order items not init")
	}
	if len(order.Products) == 0 &&
		len(order.Bundles) == 0 {
		return 0, errors.New("not purchases")
	}

	var total Money
	for _, product := range order.Products {
		total += ToMoney(product.Price)
	}

	for _, bundle := range order.Bundles {
		var medTotal Money
		for _, product := range bundle.Products {
			medTotal += ToMoney(product.Price)
		}
		total += medTotal.Multiply(bundle.Discount)
	}

	if total <= ToMoney(0.0) {
		return 0, errors.New("total cannot be negative: " + total.String())
	}
	return total.Float32(), nil

}

//Подсчет суммы заказа для клиента
func (shop Shop) PlaceOrder(username string, order shop_competition.Order) error {
	if order.Products == nil &&
		order.Bundles == nil {
		return errors.New("order items not init")
	}
	if len(order.Products) == 0 &&
		len(order.Bundles) == 0 {
		return errors.New("not purchases")
	}

	user, ok := shop.Accounts[username]
	if !ok {
		return errors.New("user is not registered")
	}

	total := Money(0)
	//Products
	if len(order.Products) != 0 {
		totalProducts := Money(0)
		cacheProducts := shop.getCacheProductsTotal(order, user.AccountType)
		if cacheProducts.IsHas {
			totalProducts = cacheProducts.Total
		} else {
			for _, product := range order.Products {
				totalProducts += ToMoney(product.Price).Multiply(getCoefficientByType(user.AccountType, product.Type))
			}
			shop.CacheProducts[cacheProducts.Key] = totalProducts
		}
		total += totalProducts
	}
	//Bundles
	if len(order.Bundles) != 0 {
		totalBundles := Money(0)
		//cacheBundles := shop.getCacheBundlesTotal(order, user.AccountType)
		//if cacheBundles.IsHas{
		//totalBundles = cacheBundles.Total
		//} else {
		for _, bundle := range order.Bundles {
			merTotal := Money(0)
			for _, product := range bundle.Products {
				merTotal += ToMoney(product.Price).Multiply(getCoefficientByType(user.AccountType, product.Type))
			}
			totalBundles += merTotal.Multiply(bundle.Discount)
		}
		//shop.CacheBundles[cacheBundles.Key] = totalBundles
		//}
		total += totalBundles
	}
	if ToMoney(user.Balance) < total {
		return errors.New("user has insufficient balance")
	}

	if total < ToMoney(0.0) {
		return errors.New("total cannot be negative: " + total.String())
	}
	user.Balance -= total.Float32()
	shop.Accounts[username] = user
	return nil
}

// Проверка наличия кэшированного значения суммы заказа для продуктов
// Формат кэша: СтатусКлиентаНаименование1Наименование2...
func (shop Shop) getCacheProductsTotal(order shop_competition.Order, accountType shop_competition.AccountType) CacheInfo {
	// Сортируем продукты в правильный набор
	sortProducts := make([]shop_competition.Product, len(order.Products))
	copy(sortProducts, order.Products)
	sort.Slice(sortProducts, func(i, j int) bool {
		return sortProducts[i].Name < sortProducts[j].Name
	})

	orderKey := string(accountType)
	for _, v := range sortProducts {
		orderKey += v.Name
	}
	//for _, bundle := range order.Bundles{
	//	orderKey += "{"
	//	for _, product := range bundle.Products{
	//		orderKey += product.Name
	//	}
	//	orderKey += "}"
	//}

	total, ok := shop.CacheProducts[orderKey]

	return CacheInfo{
		IsHas: ok,
		Total: total,
		Key:   orderKey,
	}
}

// Проверка наличия кэшированного значения суммы заказа для продуктов
// Формат кэша: СтатусКлиента{Наименование1Наименование2...}...
//func (shop Shop) getCacheBundlesTotal(order shop_competition.Order, accountType shop_competition.AccountType) CacheInfo {
// Сортируем продукты в правильный набор
//sortProducts := make([]shop_competition.Bundle, len(order.Bundles))
//copy(sortProducts, order.Products)
//sort.Slice(sortProducts, func(i, j int) bool {
//	return sortProducts[i].Name < sortProducts[j].Name
//})
//
//orderKey := string(accountType)
////for _, bundle := range order.Bundles{
////	orderKey += "{"
////	for _, product := range bundle.Products{
////		orderKey += product.Name
////	}
////	orderKey += "}"
////}
//
//total, ok := shop.CacheProducts[orderKey]

//return CacheInfo{
//	IsHas: ok,
//	Total: total,
//	Key:   orderKey,
//}
//}
func getCoefficientByType(accountType shop_competition.AccountType, productType shop_competition.ProductType) float32 {
	var coefficient float32 = 1.0
	switch {
	case accountType == shop_competition.AccountNormal &&
		productType == shop_competition.ProductPremium:
		coefficient = 1.5
	case accountType == shop_competition.AccountPremium &&
		productType == shop_competition.ProductNormal:
		coefficient = 0.95
	case accountType == shop_competition.AccountPremium &&
		productType == shop_competition.ProductPremium:
		coefficient = 0.80
	}
	return coefficient
}

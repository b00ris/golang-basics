package shop

import (
	"encoding/json"
	"errors"
	"github.com/youricorocks/shop_competition"
)

var cacheErr = struct {
	ReCalc bool
	err    error
}{}

//Подсчет суммы заказа
func (shop Shop) CalculateOrder(username string, order shop_competition.Order) (float32, error) {
	if order.Products == nil &&
		order.Bundles == nil {
		return 0, errors.New("order items not init")
	}
	if len(order.Products) == 0 &&
		len(order.Bundles) == 0 {
		return 0, errors.New("not purchases")
	}

	var total Money

	//If send accountType aka params this check not necessary
	user, ok := shop.Accounts[username]
	if !ok {
		return 0, errors.New("user is not registered")
	}
	var cacheKey []byte
	cacheKey, cacheErr.err = json.Marshal(order)
	if cacheErr.err == nil {
		totalCache, ok := shop.CacheProducts[string(cacheKey)]
		if ok {
			return totalCache.Float32(), nil
		}
	} else {
		cacheErr.ReCalc = true
	}
	//Products
	for i := 0; i < len(order.Products); i++ {
		product := order.Products[i]
		if product.Type == shop_competition.ProductSample {
			cacheErr.ReCalc = true
			order.Products = append(order.Products[:i], order.Products[i+1:]...)
			i--
		}
		total += ToMoney(product.Price).Multiply(getCoefficientByType(user.AccountType, product.Type))
	}

	//Bundles
	for _, bundle := range order.Bundles {
		merTotal := Money(0)
		for _, product := range bundle.Products {
			merTotal += ToMoney(product.Price).Multiply(getCoefficientByType(user.AccountType, product.Type))
		}
		total += merTotal.Multiply(bundle.Discount)
	}

	if cacheErr.ReCalc {
		cacheKey, cacheErr.err = json.Marshal(order)
	}
	if cacheErr.err == nil {
		shop.CacheProducts[string(cacheKey)] = total
	}
	if total <= ToMoney(0.0) {
		return 0, errors.New("total cannot be negative: " + total.String())
	}
	return total.Float32(), nil
}

//Подсчет суммы заказа для клиента
func (shop Shop) PlaceOrder(username string, order shop_competition.Order) error {
	user, ok := shop.Accounts[username]
	if !ok {
		return errors.New("user is not registered")
	}

	total, err := shop.CalculateOrder(username, order)
	if err != nil {
		return err
	}

	if ToMoney(user.Balance) < ToMoney(total) {
		return errors.New("user has insufficient balance")
	}

	user.Balance -= total
	shop.Accounts[username] = user
	return nil
}

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

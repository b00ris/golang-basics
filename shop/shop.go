package shop

import (
	"bytes"

	"encoding/json"
	"errors"
	"github.com/youricorocks/shop_competition"

	"sync"
	"time"
)

type Shop struct {
	Products
	Bundles
	Accounts
	CacheProducts map[string]Money
	sync.RWMutex
}

func NewShop() *Shop {
	return &Shop{
		Products:      Products{Products: make(map[string]shop_competition.Product)},
		Bundles:       Bundles{Bundles: make(map[string]shop_competition.Bundle)},
		Accounts:      Accounts{Accounts: make(map[string]shop_competition.Account)},
		CacheProducts: make(map[string]Money),
	}
}

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
	user, ok := shop.Accounts.Accounts[username]
	if !ok {
		return 0, errors.New("user is not registered")
	}
	var cacheKey []byte
	var cacheErr = struct {
		ReCalc bool
		err    error
	}{}
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
		shop.Lock()
		shop.CacheProducts[string(cacheKey)] = total
		shop.Unlock()
	}
	if total <= ToMoney(0) {
		return 0, errors.New("total cannot be negative: " + total.String())
	}
	return total.Float32(), nil
}

func (shop Shop) PlaceOrderConc(username string, order shop_competition.Order) error {
	ch := make(chan error, 1)
	go func() {
		ch <- shop.PlaceOrder(username, order)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}

//Подсчет суммы заказа для клиента
func (shop Shop) PlaceOrder(username string, order shop_competition.Order) error {
	user, ok := shop.Accounts.Accounts[username]
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
	shop.Accounts.Lock()
	defer shop.Accounts.Unlock()

	shop.Accounts.Accounts[username] = user

	return nil
}

func getCoefficientByType(accountType shop_competition.AccountType, productType shop_competition.ProductType) float32 {
	var coefficient float32 = 1.0

	if accountType == shop_competition.AccountNormal && productType == shop_competition.ProductPremium {
		coefficient = 1.5
	} else if accountType == shop_competition.AccountPremium && productType == shop_competition.ProductNormal {
		coefficient = 0.95
	} else if accountType == shop_competition.AccountPremium && productType == shop_competition.ProductPremium {
		coefficient = 0.80
	}
	return coefficient
}

func (shop Shop) Import(data []byte) error {
	return json.Unmarshal(data, &shop)
}

func (shop Shop) Export() ([]byte, error) {
	reqBodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(shop); err != nil {
		return nil, err
	}

	return reqBodyBytes.Bytes(), nil
}

const PACKAGE_SIZE = 1000
const (
	ImportNone = iota
	ImportProcessing
	ImportEof
)

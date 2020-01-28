package shop

import "github.com/youricorocks/shop_competition"

type Products map[string]shop_competition.Product
type Bundles map[string]shop_competition.Bundle
type Accounts map[string]shop_competition.Account
type Shop struct {
	Products
	Bundles
	Accounts
	CacheProducts map[string]Money
}

func ShopInit() *Shop {
	return &Shop{
		Accounts:      make(map[string]shop_competition.Account),
		CacheProducts: make(map[string]Money),
	}
}

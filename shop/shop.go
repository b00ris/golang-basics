package shop

import "github.com/youricorocks/shop_competition"

type Shop struct {
	Accounts      map[string]shop_competition.Account
	CacheProducts map[string]Money
}

func ShopInit() *Shop {
	return &Shop{
		Accounts:      make(map[string]shop_competition.Account),
		CacheProducts: make(map[string]Money),
	}
}

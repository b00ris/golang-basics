package test

import (
	"github.com/youricorocks/shop_competition"
	"lessons/basics/shop"
	"testing"
)

func TestSort(t *testing.T) {
	shopTest := shop.ShopInit()
	shopTest.Accounts = map[string]shop_competition.Account{
		"Bred":   {Name: "Bred", Balance: 1000, AccountType: shop_competition.AccountNormal},
		"Alfrad": {Name: "Alfrad", Balance: 900, AccountType: shop_competition.AccountNormal},
	}

	accs := shopTest.Accounts.GetAccounts(shop_competition.SortByBalance)
	for _, v := range accs {
		t.Log(v)
	}
}

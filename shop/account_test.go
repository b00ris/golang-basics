package shop

import (
	"github.com/youricorocks/shop_competition"
	"sync"
	"testing"
)

func TestSort(t *testing.T) {
	shopTest := NewShop()
	shopTest.Accounts = Accounts{
		Accounts: map[string]shop_competition.Account{
			"Bred":   {Name: "Bred", Balance: 1000, AccountType: shop_competition.AccountNormal},
			"Alfrad": {Name: "Alfrad", Balance: 900, AccountType: shop_competition.AccountNormal},
		},
		RWMutex: sync.RWMutex{},
	}

	accs := shopTest.Accounts.GetAccounts(shop_competition.SortByBalance)
	for _, v := range accs {
		t.Log(v)
	}
}

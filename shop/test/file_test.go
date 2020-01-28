package test

import (
	//"github.com/youricorocks/shop_competition"
	"github.com/youricorocks/shop_competition"
	"lessons/basics/shop"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	shop1 := NewShopAccounts(map[string]shop_competition.Account{"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}})
	shopBytes, err := shop1.Export()
	if err != nil {
		t.Fatal(err)
	}
	shop2 := shop.NewShop()
	err = shop2.Import(shopBytes)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(shop1, shop2) {
		t.Fatal()
	}
}

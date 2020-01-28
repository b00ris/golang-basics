package test

import (
	//"github.com/youricorocks/shop_competition"
	"github.com/youricorocks/shop_competition"
	shop3 "lessons/basics/shop"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	shop := shopAccountsInit(map[string]shop_competition.Account{"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}})
	shopBytes, err := shop.Export()
	if err != nil {
		t.Fatal(err)
	}
	shop2 := shop3.ShopInit()
	err = shop2.Import(shopBytes)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(shop, shop2) {
		t.Fatal()
	}
}

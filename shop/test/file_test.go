package test

import (
	//"github.com/youricorocks/shop_competition"
	"github.com/youricorocks/shop_competition"
	shop3 "lessons/basics/shop"
	"testing"
)

func TestImport(t *testing.T) {
	shop := shopAccountsInit(map[string]shop_competition.Account{"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}})
	byte, _ := shop.Export()
	shop2 := shop3.ShopInit()
	shop2.Import(byte)
}

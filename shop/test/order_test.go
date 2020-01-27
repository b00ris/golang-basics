package test

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	"lessons/basics/shop"
	"testing"
)

// {CalculateOrderBlock}
type OrderTest struct {
	shop     *shop.Shop
	order    shop_competition.Order
	username string
	result   float32
	err      error
}

var calculateOrderTests = []OrderTest{
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		result: 200,
		err:    nil,
	},
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{},
			Bundles: []shop_competition.Bundle{
				{
					Products: []shop_competition.Product{
						{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
						{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
					},
					Type:     shop_competition.BundleNormal,
					Discount: 0.5,
				},
			},
		},
		result: 100,
		err:    nil,
	},
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
			},
			Bundles: []shop_competition.Bundle{
				{
					Products: []shop_competition.Product{
						{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
						{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
					},
					Type:     shop_competition.BundleNormal,
					Discount: 0.5,
				},
			},
		},
		result: 300,
		err:    nil,
	},
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 0.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: -1.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		result: 0,
		err:    errors.New("total cannot be negative: -0.99"),
	},
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: nil,
			Bundles:  nil,
		},
		result: 0,
		err:    errors.New("order items not init"),
	},
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{},
			Bundles:  []shop_competition.Bundle{},
		},
		result: 0,
		err:    errors.New("not purchases"),
	},
}

func TestCalculateOrder(t *testing.T) {
	for _, v := range calculateOrderTests {
		total, err := v.shop.CalculateOrder(v.order)

		if v.err == nil {
			if err != nil {
				t.Fatal(err)
			}
			if total != v.result {
				t.Fatal("Price does not match expected: ", total, " != ", v.result)
			}
		} else {
			if err == nil {
				t.Fatal("Error expected: ", v.err, ", but get total = ", total)
			}
			if err.Error() != v.err.Error() {
				t.Fatal("Error is not correct:", err, ", wait: ", v.err)
			}
		}
	}
}

// {CalculateOrderBlock}

// {PlaceOrderBlock}
func shopAccountsInit(accounts map[string]shop_competition.Account) *shop.Shop {
	return &shop.Shop{
		Accounts:      accounts,
		CacheProducts: make(map[string]shop.Money),
	}
}

var placeOrderTests = []OrderTest{
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		username: "BredFalcon",
		result:   800,
		err:      nil,
	}, // Products
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{},
			Bundles: []shop_competition.Bundle{
				{
					Products: []shop_competition.Product{
						{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
						{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
					},
					Type:     shop_competition.BundleNormal,
					Discount: 0.5,
				},
			},
		},
		username: "BredFalcon",
		result:   900,
		err:      nil,
	}, // Bundles
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 100.00, Type: shop_competition.ProductNormal},
			},
			Bundles: []shop_competition.Bundle{
				{
					Products: []shop_competition.Product{
						{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
						{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
					},
					Type:     shop_competition.BundleNormal,
					Discount: 0.5,
				},
			},
		},
		username: "BredFalcon",
		result:   800,
		err:      nil,
	}, // Combination
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 100.00, Type: shop_competition.ProductPremium},
				{Name: "Potato", Price: 33.33, Type: shop_competition.ProductNormal},
			},
			Bundles: []shop_competition.Bundle{},
		},
		username: "BredFalcon",
		result:   816.67,
		err:      nil,
	}, // PremiumItem
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal},
			"KelvinKlay": {Name: "Kelvin", Balance: 1000.0, AccountType: shop_competition.AccountPremium}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 100.00, Type: shop_competition.ProductPremium},
				{Name: "Potato", Price: 100.00, Type: shop_competition.ProductNormal},
			},
			Bundles: []shop_competition.Bundle{
				{
					Products: []shop_competition.Product{
						{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
						{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
					},
					Type:     shop_competition.BundleNormal,
					Discount: 0.5,
				},
			},
		},
		username: "KelvinKlay",
		result:   730,
		err:      nil,
	}, // PremiumUser
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: 50.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		username: "Bred",
		result:   0,
		err:      errors.New("user is not registered"),
	}, // UserNotFound
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 500.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
				{Name: "Pineapple", Price: 150.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		username: "BredFalcon",
		result:   0,
		err:      errors.New("user has insufficient balance"),
	}, // BalanceInfluence
	{
		shop: shopAccountsInit(map[string]shop_competition.Account{
			"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}}),
		order: shop_competition.Order{
			Products: []shop_competition.Product{
				{Name: "Pineapple", Price: 0.00, Type: shop_competition.ProductNormal},
				{Name: "Mango", Price: -1.00, Type: shop_competition.ProductNormal},
			},
			Bundles: nil,
		},
		username: "BredFalcon",
		result:   0,
		err:      errors.New("total cannot be negative: -0.98"),
	}, // NegativeTotal
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: nil,
			Bundles:  nil,
		},
		result: 0,
		err:    errors.New("order items not init"),
	}, // NotInit
	{
		shop: shop.ShopInit(),
		order: shop_competition.Order{
			Products: []shop_competition.Product{},
			Bundles:  []shop_competition.Bundle{},
		},
		result: 0,
		err:    errors.New("not purchases"),
	}, // NotPurchase
}

func TestPlaceOrder(t *testing.T) {
	for i, v := range placeOrderTests {
		err := v.shop.PlaceOrder(v.username, v.order)

		user := v.shop.Accounts[v.username]
		if v.err == nil {
			if err != nil {
				t.Fatal(i, ". ", err)
			}
			if user.Balance != v.result {
				t.Fatal(i, ". Price does not match expected: ", user.Balance, " != ", v.result)
			}
		} else {
			if err == nil {
				t.Fatal(i, ". Error expected: ", v.err, ", but get total = ", user.Balance)
			}
			if err.Error() != v.err.Error() {
				t.Fatal(i, ". Error is not correct:", err, ", wait: ", v.err)
			}
		}
	}
}

func TestCache(t *testing.T) {
	testPlace := placeOrderTests[0]

	userBalance := testPlace.shop.Accounts[testPlace.username].Balance
	err := testPlace.shop.PlaceOrder(testPlace.username, testPlace.order)
	if err != nil {
		t.Fatal(err)
	}
	dif := userBalance - testPlace.shop.Accounts[testPlace.username].Balance
	userBalance = testPlace.shop.Accounts[testPlace.username].Balance
	err = testPlace.shop.PlaceOrder(testPlace.username, testPlace.order)
	if err != nil {
		t.Fatal(err)
	}
	difCache := userBalance - testPlace.shop.Accounts[testPlace.username].Balance
	if dif != difCache {
		t.Fatal("Cache has a different meaning")
	}
}

// {PlaceOrderBlock}

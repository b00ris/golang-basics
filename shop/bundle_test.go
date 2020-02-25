package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	"sync"
	"testing"
)

type bundleTest struct {
	shop     *Shop
	name     string
	main     shop_competition.Product
	discount float32
	adds     []shop_competition.Product
	//bundle   shop_competition.Bundle
	err error
}

func NewShopBundles(bundles map[string]shop_competition.Bundle) *Shop {
	return &Shop{
		Bundles: Bundles{
			Bundles: bundles,
			RWMutex: sync.RWMutex{},
		},
	}
}

var addBundleTests = []bundleTest{
	{
		shop: NewShopBundles(map[string]shop_competition.Bundle{}),
		name: "KingBundle",
		main: shop_competition.Product{
			Name:  "Burger King",
			Price: 250,
			Type:  shop_competition.ProductNormal,
		},
		discount: 50,
		adds: []shop_competition.Product{
			{
				Name:  "Cola",
				Price: 50,
				Type:  shop_competition.ProductNormal,
			},
		},
		err: nil,
	},
	{
		shop: NewShopBundles(map[string]shop_competition.Bundle{}),
		name: "KingBundle",
		main: shop_competition.Product{
			Name:  "Burger King",
			Price: 250,
			Type:  shop_competition.ProductNormal,
		},
		discount: 50,
		adds: []shop_competition.Product{
			{
				Name:  "Cola",
				Price: 0,
				Type:  shop_competition.ProductSample,
			},
		},
		err: nil,
	},
	{
		shop: NewShopBundles(map[string]shop_competition.Bundle{}),
		name: "KingBundle",
		main: shop_competition.Product{
			Name:  "Burger King",
			Price: 250,
			Type:  shop_competition.ProductNormal,
		},
		discount: 50,
		adds: []shop_competition.Product{
			{
				Name:  "Cola",
				Price: 50,
				Type:  shop_competition.ProductSample,
			},
		},
		err: errors.New("sample don`t has price"),
	},
	{
		shop: NewShopBundles(map[string]shop_competition.Bundle{}),
		name: "KingBundle",
		main: shop_competition.Product{
			Name:  "Burger King",
			Price: 250,
			Type:  shop_competition.ProductNormal,
		},
		discount: 200,
		adds: []shop_competition.Product{
			{
				Name:  "Cola",
				Price: 50,
				Type:  shop_competition.ProductSample,
			},
		},
		err: errors.New("discount not correct"),
	},
	{
		shop: NewShopBundles(map[string]shop_competition.Bundle{}),
		name: "KingBundle",
		main: shop_competition.Product{
			Name:  "Burger King",
			Price: 250,
			Type:  shop_competition.ProductNormal,
		},
		discount: 50,
		adds: []shop_competition.Product{
			{
				Name:  "Cola",
				Price: 50,
				Type:  shop_competition.ProductSample,
			},
			{
				Name:  "Cola",
				Price: 50,
				Type:  shop_competition.ProductSample,
			},
		},
		err: errors.New("samples bundle has only one sample"),
	},
}

func TestAddBundle(t *testing.T) {
	for i, test := range addBundleTests {
		err := test.shop.Bundles.AddBundle(test.name, test.main, test.discount, test.adds...)
		if test.err == nil {
			if err != nil {
				t.Fatal(i, ". ", err)
			}

			if _, ok := test.shop.Bundles.Bundles[test.name]; !ok {
				t.Fatal(i, ". Product not added")
			}
		} else {

			if err == nil || err.Error() != test.err.Error() {
				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
			}
		}
	}
}

//var modifyProductTests = []ProductTest{
//	{
//		shop:  shopProductsInit(map[string]shop_competition.Product{
//			"Pineapple": {Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
//		}) ,
//		product: shop_competition.Product{
//			Name:  "Pineapple",
//			Price: 500,
//			Type:  shop_competition.ProductPremium,
//		},
//		err:     nil,
//	},
//	{
//		shop:  shopProductsInit(map[string]shop_competition.Product{
//		}) ,
//		product: shop_competition.Product{
//			Name:  "Pineapple",
//			Price: 500,
//			Type:  shop_competition.ProductPremium,
//		},
//		err:     errors.New("product not found"),
//	},
//}
//func TestModifyProduct(t *testing.T) {
//	for i, test := range modifyProductTests{
//		err := test.shop.Products.ModifyProduct(test.product)
//		if test.err == nil{
//			if err != nil{
//				t.Fatal(i, ". ", err)
//			}
//
//			if product, ok := test.shop.Products[test.product.Name]; !ok || !reflect.DeepEqual(test.product, product){
//				t.Fatal(i, ". Product not modified")
//			}
//		} else {
//
//			if err == nil || err.Error() != test.err.Error() {
//				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
//			}
//		}
//	}
//}
//
//var deleteProductTests = []ProductTest{
//	{
//		shop:  shopProductsInit(map[string]shop_competition.Product{
//			"Pineapple": {Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
//		}) ,
//		product: shop_competition.Product{
//			Name:  "Pineapple",
//			Price: 500,
//			Type:  shop_competition.ProductPremium,
//		},
//		err:     nil,
//	},
//	{
//		shop:  shopProductsInit(map[string]shop_competition.Product{
//		}) ,
//		product: shop_competition.Product{
//			Name:  "Pineapple",
//			Price: 500,
//			Type:  shop_competition.ProductPremium,
//		},
//		err:     errors.New("product not found"),
//	},
//}
//func TestDeleteProduct(t *testing.T) {
//	for i, test := range deleteProductTests{
//		err := test.shop.Products.RemoveProduct(test.product.Name)
//		if test.err == nil{
//			if err != nil{
//				t.Fatal(i, ". ", err)
//			}
//
//			if _, ok := test.shop.Products[test.product.Name]; ok{
//				t.Fatal(i, ". Product not deleted")
//			}
//		} else {
//
//			if err == nil || err.Error() != test.err.Error() {
//				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
//			}
//		}
//	}
//

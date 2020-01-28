package test

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	"lessons/basics/shop"
	"reflect"
	"testing"
)

type ProductTest struct {
	shop    *shop.Shop
	product shop_competition.Product
	err     error
}

func shopProductsInit(products map[string]shop_competition.Product) *shop.Shop {
	return &shop.Shop{
		Products: products,
	}
}

var addProductTests = []ProductTest{
	{
		shop:    shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
		err:     nil,
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 1000,
			Type:  shop_competition.ProductPremium,
		},
		err: nil,
	},
	{
		shop:    shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{Name: "", Price: 1000, Type: shop_competition.ProductNormal},
		err:     errors.New("product without name"),
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{
			"Pineapple": {Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
		}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 1000,
			Type:  shop_competition.ProductPremium,
		},
		err: errors.New("product exist"),
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: -1,
			Type:  shop_competition.ProductPremium,
		},
		err: errors.New("product price -1.00 not valid"),
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 0,
			Type:  shop_competition.ProductPremium,
		},
		err: errors.New("product price 0.00 not valid"),
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 0,
			Type:  shop_competition.ProductSample,
		},
		err: nil,
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 111,
			Type:  shop_competition.ProductSample,
		},
		err: errors.New("sample was free in bundle"),
	},
}

func TestAddProduct(t *testing.T) {
	for i, test := range addProductTests {
		err := test.shop.Products.AddProduct(test.product)
		if test.err == nil {
			if err != nil {
				t.Fatal(i, ". ", err)
			}

			if _, ok := test.shop.Products[test.product.Name]; !ok {
				t.Fatal(i, ". Product not added")
			}
		} else {

			if err == nil || err.Error() != test.err.Error() {
				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
			}
		}
	}
}

var modifyProductTests = []ProductTest{
	{
		shop: shopProductsInit(map[string]shop_competition.Product{
			"Pineapple": {Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
		}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 500,
			Type:  shop_competition.ProductPremium,
		},
		err: nil,
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 500,
			Type:  shop_competition.ProductPremium,
		},
		err: errors.New("product not found"),
	},
}

func TestModifyProduct(t *testing.T) {
	for i, test := range modifyProductTests {
		err := test.shop.Products.ModifyProduct(test.product)
		if test.err == nil {
			if err != nil {
				t.Fatal(i, ". ", err)
			}

			if product, ok := test.shop.Products[test.product.Name]; !ok || !reflect.DeepEqual(test.product, product) {
				t.Fatal(i, ". Product not modified")
			}
		} else {

			if err == nil || err.Error() != test.err.Error() {
				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
			}
		}
	}
}

var deleteProductTests = []ProductTest{
	{
		shop: shopProductsInit(map[string]shop_competition.Product{
			"Pineapple": {Name: "Pineapple", Price: 1000, Type: shop_competition.ProductNormal},
		}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 500,
			Type:  shop_competition.ProductPremium,
		},
		err: nil,
	},
	{
		shop: shopProductsInit(map[string]shop_competition.Product{}),
		product: shop_competition.Product{
			Name:  "Pineapple",
			Price: 500,
			Type:  shop_competition.ProductPremium,
		},
		err: errors.New("product not found"),
	},
}

func TestDeleteProduct(t *testing.T) {
	for i, test := range deleteProductTests {
		err := test.shop.Products.RemoveProduct(test.product.Name)
		if test.err == nil {
			if err != nil {
				t.Fatal(i, ". ", err)
			}

			if _, ok := test.shop.Products[test.product.Name]; ok {
				t.Fatal(i, ". Product not deleted")
			}
		} else {

			if err == nil || err.Error() != test.err.Error() {
				t.Fatal(i, ". Error is not correct:", err, ", wait: ", test.err)
			}
		}
	}
}

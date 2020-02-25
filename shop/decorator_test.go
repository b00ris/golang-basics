package shop

import (
	"github.com/youricorocks/shop_competition"
	"testing"
	"time"
)


func TestAddProductConcurrencySuccess(t *testing.T) {
	tests := []struct {
		name     string
		Products []shop_competition.Product
		duration time.Duration
	}{
		{
			name: "A",
			Products: []shop_competition.Product{
				{Name: "A", Price: 100, Type: shop_competition.ProductNormal},
				{Name: "B", Price: 100, Type: shop_competition.ProductNormal},
				{Name: "C", Price: 100, Type: shop_competition.ProductNormal},
			},
			duration: time.Second,
		},
	}

	testShopOrig := NewShop()
	testShop:= TimeoutDecorator{
		shop:    testShopOrig,
		timeout: time.Second*2,
	}

	for _, test := range tests {
		test:=test
		t.Run(test.name, func(t *testing.T) {
			ch := make(chan error, len(test.Products))
			for _, product := range test.Products {
				product := product
				go func() {
					ch <- testShop.AddProduct(product)
				}()
			}

			i := len(test.Products)
			for res := range ch {
				t.Logf("%v. %v", i, res)
				i--
				if i == 0 {
					break
				}
			}
		})
	}
}

func TestName(t *testing.T) {
	td:=TimeoutDecorator{
		timeout:time.Second,
		shop: &ShopDecoratorStub{
			addProductFunc: func(product shop_competition.Product) error {
				time.Sleep(time.Second*5)
				return nil
			},
		},
	}
	err:=td.AddProduct(shop_competition.Product{})
	if err==nil {
		t.Fatal()
	}
}
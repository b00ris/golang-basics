// +build !race

package shop

import (
	"github.com/youricorocks/shop_competition"
	"testing"
	"time"
)

func TestAddProductConcurrencySuccessful1(t *testing.T) {
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

	testShop := NewShop()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ch := make(chan error, len(test.Products))
			for _, product := range test.Products {
				product := product
				go func() {
					ch <- testShop.AddProductConc(product)
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

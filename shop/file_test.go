// +build !race

package shop

import (
	"fmt"
	"github.com/youricorocks/shop_competition"
	"reflect"
	"strconv"
	"testing"
)

func TestSave(t *testing.T) {
	shop1 := NewShopAccounts(map[string]shop_competition.Account{"BredFalcon": {Name: "Bred", Balance: 1000.0, AccountType: shop_competition.AccountNormal}})
	shopBytes, err := shop1.Export()
	if err != nil {
		t.Fatal(err)
	}
	shop2 := NewShop()
	err = shop2.Import(shopBytes)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(shop1, shop2) {
		t.Fatal()
	}
}

func TestExportProductsCSV(t *testing.T) {
	testShop := NewShopProducts(map[string]shop_competition.Product{
		"A": {Name: "A", Price: 1000, Type: shop_competition.ProductNormal},
		"B": {Name: "B", Price: 500, Type: shop_competition.ProductNormal},
		"C": {Name: "C", Price: 500, Type: shop_competition.ProductNormal},
	})
	buf := testShop.ExportProductsCSV()
	t.Log(buf)
}

func TestImportProductsCSV(t *testing.T) {
	testShop := NewShop()
	testShop2 := NewShop()
	for i := 0; i < 5500; i++ {
		name := fmt.Sprintf("%v", i)
		testShop.Products.Products[name] = shop_competition.Product{
			Name:  name,
			Price: float32(i),
			Type:  shop_competition.ProductNormal,
		}
	}
	buf := testShop.Products.ExportProductsCSV()
	err := testShop2.Products.ImportProductsCSV(buf)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testShop.Products.Products, testShop2.Products.Products) {
		t.Fatal("Not Equal shops")
	}
}

func BenchmarkImportProductsCSV(b *testing.B) {
	testShop := NewShop()
	testShop2 := NewShop()
	for i := 0; i < 100000; i++ {
		name := fmt.Sprintf("%v", i)
		testShop.Products.Products[name] = shop_competition.Product{
			Name:  name,
			Price: float32(i),
			Type:  shop_competition.ProductNormal,
		}
	}
	buf := testShop.Products.ExportProductsCSV()
	b.ResetTimer()
	testShop2.Products.ImportProductsCSV(buf)
}
func TestStrconv(t *testing.T) {
	t.Log(strconv.ParseInt("4", 0, strconv.IntSize))
}

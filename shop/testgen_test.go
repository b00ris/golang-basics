package shop

import (
	"github.com/youricorocks/shop_competition"
)

type ShopDecoratorStub struct{
	addProductFunc func(shop_competition.Product) error
}

func (c ShopDecoratorStub) AddProduct(products shop_competition.Product) error {
	if c.addProductFunc !=nil {
		return c.addProductFunc(products)
	}
	return nil
}

func (ShopDecoratorStub) ModifyProduct(shop_competition.Product) error {
	panic("implement me")
}

func (ShopDecoratorStub) RemoveProduct(name string) error {
	panic("implement me")
}

func (ShopDecoratorStub) Register(username string) error {
	panic("implement me")
}

func (ShopDecoratorStub) AddBalance(username string, sum float32) error {
	panic("implement me")
}

func (ShopDecoratorStub) Balance(username string) (float32, error) {
	panic("implement me")
}

func (ShopDecoratorStub) GetAccounts(sort shop_competition.AccountSortType) []shop_competition.Account {
	panic("implement me")
}

func (ShopDecoratorStub) PlaceOrder(username string, order shop_competition.Order) error {
	panic("implement me")
}

func (ShopDecoratorStub) AddBundle(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	panic("implement me")
}

func (ShopDecoratorStub) ChangeDiscount(name string, discount float32) error {
	panic("implement me")
}

func (ShopDecoratorStub) RemoveBundle(name string) error {
	panic("implement me")
}

func (ShopDecoratorStub) Import(data []byte) error {
	panic("implement me")
}

func (ShopDecoratorStub) Export() ([]byte, error) {
	panic("implement me")
}

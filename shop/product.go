package shop

import "github.com/youricorocks/shop_competition"

type MyProduct shop_competition.Product

func (order MyOrder) AddProduct(shop_competition.Product) error {
	panic("implement me")
}

func (order MyOrder) ModifyProduct(shop_competition.Product) error {
	panic("implement me")
}

func (order MyOrder) RemoveProduct(name string) error {
	panic("implement me")
}

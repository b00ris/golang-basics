package shop

import "github.com/youricorocks/shop_competition"

type MyBundle shop_competition.Bundle

func (shop Shop) AddBundle(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	panic("implement me")
}

func (shop Shop) ChangeDiscount(name string, discount float32) error {
	panic("implement me")
}

func (shop Shop) RemoveBundle(name string) error {
	panic("implement me")
}

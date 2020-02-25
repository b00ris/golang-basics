package shop

import (
	"github.com/youricorocks/shop_competition"
	"time"
)
var _ shop_competition.Shop = &TimeoutDecorator{}

type TimeoutDecorator struct{
	shop shop_competition.Shop
	timeout time.Duration
}

//func (td *TimeoutDecorator) AddProduct(product shop_competition.Product) error {
//	ch := make(chan error, 1)
//	go func() {
//		ch <- td.shop.AddProduct(product)
//	}()
//	select {
//	case err := <-ch:
//		return err
//	case <-time.After(DURATION):
//		return TimeoutError
//	}
//}



func (td *TimeoutDecorator) ModifyProduct(product shop_competition.Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- td.shop.ModifyProduct(product)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(td.timeout):
		return TimeoutError
	}
}

func (td *TimeoutDecorator) RemoveProduct(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- td.shop.RemoveProduct(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(td.timeout):
		return TimeoutError
	}
}

func (td *TimeoutDecorator) Register(username string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- td.shop.Register(username)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(td.timeout):
		return TimeoutError
	}
}

func (td *TimeoutDecorator) AddBalance(username string, sum float32) error {
	ch := make(chan error, 1)
	go func() {
		ch <- td.shop.AddBalance(username, sum)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(td.timeout):
		return TimeoutError
	}
}

func (td *TimeoutDecorator) Balance(username string) (float32, error) {
	panic("implement me")
}

func (td *TimeoutDecorator) GetAccounts(sort shop_competition.AccountSortType) []shop_competition.Account {
	panic("implement me")
}

func (td *TimeoutDecorator) PlaceOrder(username string, order shop_competition.Order) error {
	ch := make(chan error, 1)
	go func() {
		ch <- td.shop.PlaceOrder(username, order)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(td.timeout):
		return TimeoutError
	}
}

func (td *TimeoutDecorator) AddBundle(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	return td.timeoutFunc(func(ch chan error) {
		ch <- td.shop.AddBundle(name, main, discount, additional...)
	})
}

func (td *TimeoutDecorator) ChangeDiscount(name string, discount float32) error {
	return td.timeoutFunc(func(ch chan error) {
		ch <- td.shop.ChangeDiscount(name, discount)
	})
}

func (td *TimeoutDecorator) RemoveBundle(name string) error {
	return td.timeoutFunc(func(ch chan error) {
		ch <- td.shop.RemoveBundle(name)
	})
}

func (td *TimeoutDecorator) Import(data []byte) error {
	return td.timeoutFunc(func(ch chan error) {
		ch <- td.shop.Import(data)
	})
}

func (td *TimeoutDecorator) Export() ([]byte, error) {
	panic("implement me")
}


func (td *TimeoutDecorator) AddProduct(product shop_competition.Product) error {
	return td.timeoutFunc(func(ch chan error) {
		ch <- td.shop.AddProduct(product)
	})
}
func (td *TimeoutDecorator) timeoutFunc(f func(ch chan error)) error {
	ch := make(chan error, 1)
	go func() {
		f(ch)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}
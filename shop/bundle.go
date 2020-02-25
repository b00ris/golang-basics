package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
	"sync"
	"time"
)

const OneHundredPercent = 1

type Bundles struct {
	Bundles map[string]shop_competition.Bundle
	sync.RWMutex
}

const OneItem = 1

func (bundles *Bundles) AddBundleConc(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	ch := make(chan error, 1)
	go func() {
		ch <- bundles.AddBundle(name, main, discount, additional...)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}

// discount has value in 0...100
func (bundles *Bundles) AddBundle(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	if discount < 0 || discount > 100 {
		return errors.New("discount not correct")
	}
	if len(additional) == 0 {
		return errors.New("additional products is empty")
	}
	var bundle shop_competition.Bundle

	bundle.Discount = ToMoney(OneHundredPercent - (discount / 100)).Float32()

	bundle.Products = append(additional, main)
	for _, v := range additional {
		if v.Type == shop_competition.ProductSample {
			if len(additional) != OneItem {
				return errors.New("samples bundle has only one sample")
			}
			if v.Price != 0 {
				return errors.New("sample don`t has price")
			}
			bundle.Type = shop_competition.BundleSample
		}
	}
	bundles.Lock()
	defer bundles.Unlock()
	bundle, ok := bundles.Bundles[name]
	if ok {
		return errors.New("bundle exist")
	}

	bundles.Bundles[name] = bundle
	return nil

}

func (bundles *Bundles) ChangeDiscountConc(name string, discount float32) error {
	ch := make(chan error, 1)
	go func() {
		ch <- bundles.ChangeDiscount(name, discount)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}

func (bundles *Bundles) ChangeDiscount(name string, discount float32) error {
	if discount < 0 || discount > 100 {
		return errors.New("discount not correct")
	}
	discountMoney := Money(OneHundredPercent - (discount / 100))

	bundles.Lock()
	defer bundles.Unlock()
	bundle, ok := bundles.Bundles[name]
	if !ok {
		return errors.New("bundle not found")
	}

	bundle.Discount = discountMoney.Float32()

	bundles.Bundles[name] = bundle
	return nil
}

func (bundles *Bundles) RemoveBundleConc(name string) error {
	ch := make(chan error, 1)
	go func() {
		ch <- bundles.RemoveBundle(name)
	}()
	select {
	case err := <-ch:
		return err
	case <-time.After(DURATION):
		return TimeoutError
	}
}

func (bundles *Bundles) RemoveBundle(name string) error {
	bundles.Lock()
	defer bundles.Unlock()

	if _, ok := bundles.Bundles[name]; !ok {
		return errors.New("bundle not found")
	}
	delete(bundles.Bundles, name)
	return nil
}

package shop

import (
	"errors"
	"github.com/youricorocks/shop_competition"
)

const OneHundredPercent = 1

// discount has value in 0...100
func (bundles *Bundles) AddBundle(name string, main shop_competition.Product, discount float32, additional ...shop_competition.Product) error {
	bundle, ok := (*bundles)[name]
	if ok {
		return errors.New("bundle exist")
	}
	if discount < 0 || discount > 100 {
		return errors.New("discount not correct")
	}
	bundle.Discount = OneHundredPercent - (discount / 100)

	if len(additional) == 0 {
		return errors.New("additional products is empty")
	}

	bundle.Products = append(additional, main)
	for _, v := range additional {
		if v.Type == shop_competition.ProductSample {
			if len(additional) != 1 {
				return errors.New("samples bundle has only one sample")
			}
			if v.Price != 0.0 {
				return errors.New("sample don`t has price")
			}
			bundle.Type = shop_competition.BundleSample
		}
	}

	(*bundles)[name] = bundle
	return nil

}

func (bundles *Bundles) ChangeDiscount(name string, discount float32) error {
	bundle, ok := (*bundles)[name]
	if !ok {
		return errors.New("bundle not found")
	}
	if discount < 0 || discount > 100 {
		return errors.New("discount not correct")
	}

	bundle.Discount = OneHundredPercent - (discount / 100)

	return nil
}

func (bundles *Bundles) RemoveBundle(name string) error {
	if _, ok := (*bundles)[name]; !ok {
		return errors.New("bundle not found")
	}
	delete(*bundles, name)
	return nil
}

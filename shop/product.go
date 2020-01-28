package shop

import (
	"errors"
	"fmt"
	"github.com/youricorocks/shop_competition"
)

func (products *Products) AddProduct(product shop_competition.Product) error {
	if len(product.Name) == 0 {
		return errors.New("product without name")
	}
	if _, ok := (*products)[product.Name]; ok {
		return errors.New("product exist")
	}

	err := productCheck(product)
	if err != nil {
		return err
	}

	(*products)[product.Name] = product

	return nil
}

func (products *Products) ModifyProduct(product shop_competition.Product) error {
	if _, ok := (*products)[product.Name]; !ok {
		return errors.New("product not found")
	}

	err := productCheck(product)
	if err != nil {
		return err
	}
	(*products)[product.Name] = product

	return nil
}

func (products *Products) RemoveProduct(name string) error {
	if _, ok := (*products)[name]; !ok {
		return errors.New("product not found")
	}
	delete(*products, name)
	return nil
}

func productCheck(product shop_competition.Product) error {
	if product.Type == shop_competition.ProductSample {
		if product.Price != 0 {
			return errors.New("sample was free in bundle")
		}
	} else {
		if product.Price <= 0.0 {
			return fmt.Errorf("product price %.2f not valid", product.Price)
		}
	}
	return nil
}

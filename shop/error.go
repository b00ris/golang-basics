package shop

import (
	"fmt"
	"github.com/youricorocks/shop_competition"
)

type ImportProductError struct {
	Product shop_competition.Product
	Err     error
}

func (e *ImportProductError) Error() string {
	return fmt.Sprintf("import error: product %v, err: %v", e.Product, e.Err)
}

type ImportAccountError struct {
	Account shop_competition.Account
	Err     error
}

func (e *ImportAccountError) Error() string {
	return fmt.Sprintf("import error: product %v, err: %v", e.Account, e.Err)
}

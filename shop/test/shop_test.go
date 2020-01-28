package test

import (
	"lessons/basics/shop"
	"testing"
)

func TestCalcOrder(t *testing.T) {
	var total shop.Money
	var medTotal shop.Money = shop.ToMoney(100)
	total += medTotal.Multiply(0.5)
	//total += medTotal - medTotal.Multiply(1.0/100.0) * medTotal.Multiply(50)
	t.Log(total)
}

func TestName2(t *testing.T) {
	t.Log(shop.ToMoney(1.927))
}

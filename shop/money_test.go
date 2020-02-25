package shop

import (
	"testing"
)

type moneyTest struct {
	num         float32
	coefficient float32
	res         float32
}

var toMoneyTests = []moneyTest{
	{num: 0, res: 0},
	{num: -1, res: -1},
	{num: 0.99, res: 0.99},
	{num: 0.999, res: 1},
	{num: 100.992, res: 100.99},
}

func TestToMoney(t *testing.T) {
	for i, v := range toMoneyTests {
		res := ToMoney(v.num)
		if res.Float32() != v.res {
			t.Fatal(i, ". Nums not equal:", res, " != ", v.res)
		}
	}
}

var multiplyTests = []moneyTest{
	{num: 0, coefficient: 100, res: 0},
	{num: -1, coefficient: 1, res: -1},
	{num: 0.99, coefficient: -1, res: -0.99},
	{num: 0.999, coefficient: 2, res: 2},
	{num: 100.992, coefficient: 2, res: 201.98},
}

func TestMultiply(t *testing.T) {
	for i, v := range multiplyTests {
		res := ToMoney(v.num)
		res = res.Multiply(v.coefficient)
		if res.Float32() != v.res {
			t.Fatal(i, ". Nums not equal:", res, " != ", v.res)
		}
	}
}

package shop

import "fmt"

type Money int64

const penny = 100

// ToMoney конвертирует float32 в Money
func ToMoney(f float32) Money {
	var rounding float32 = 0.5
	if f < 0 {
		rounding = -rounding
	}

	return Money((f * 100) + rounding)
}

// Float32 конвертирует Money в float32
func (m Money) Float32() float32 {
	x := float32(m)
	x = x / penny
	return x
}

// Безопасное умножение Money на float64, с округлением
func (m Money) Multiply(f float32) Money {
	var rounding float32 = 0.5

	x := float32(m) * f
	if x < 0 {
		rounding = -rounding
	}

	x += rounding

	return Money(x)
}

func (m Money) String() string {
	x := float32(m)
	x = x / penny
	return fmt.Sprintf("%.2f", x)
}

package shop

import "fmt"

type Money int64

// ToMoney конвертирует float32 в Money
func ToMoney(f float32) Money {
	return Money((f * 100) + 0.5)
}

// Float32 конвертирует Money в float32
func (m Money) Float32() float32 {
	x := float32(m)
	x = x / 100
	return x
}

// Безопасное умножение Money на float64, с округлением
func (m Money) Multiply(f float32) Money {
	x := (float32(m) * f) + 0.5
	return Money(x)
}

func (m Money) String() string {
	x := float32(m)
	x = x / 100
	return fmt.Sprintf("%.2f", x)
}

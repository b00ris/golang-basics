package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
	//wc.Test(wordCount)
	//2. Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
	var arr = make([]int, 1000)
	for i := 0; i < 1000; i++ {
		arr[i] = rand.Intn(2)
	}
	//fmt.Println(arr)
	m := intCount(arr)
	fmt.Println(m)

	//4.Сделать Фибоначчи с мемоизацией
	res = append(res, 0)
	res = append(res, 1)
	print(fib(9))
}

//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
func wordCount(str string) map[string]int {
	m := make(map[string]int)
	arr := strings.Split(str, " ")
	for _, v := range arr {
		elem, okay := m[v]
		if okay == false {
			m[v] = 1
		} else {
			m[v] += elem
		}
	}
	return m
}

//2. Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func intCount(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		_, okay := m[v]
		if okay == false {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	return m
}

var res []int

func fib(n int) int {
	if n == 0 {
		return res[0]
	}
	if n == 1 {
		return res[1]
	}
	if len(res) <= n {
		res = append(res, fib(n-2)+fib(n-1))
	}
	return res[n]
}

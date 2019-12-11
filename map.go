package main

import (
	"fmt"
	//"math/rand"
	"strings"
)

type available struct {
	isFirst  bool
	isSecond bool
}

func main2() {
	//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
	//wc.Test(wordCount)
	//2. Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
	var arr = make([]int, 100)
	for i := 0; i < 100; i++ {
		//arr[i] = rand.Intn(2)
		arr[i] = i
	}
	//fmt.Println(arr)
	//m := intCount(arr)
	//fmt.Println(m)
	//3.Есть два больших массива чисел, надо найти, какие упоминаются в обоих
	var arr2 = make([]int, 50)
	for i := 0; i < 50; i++ {
		//arr2[i] = rand.Intn(1)
		arr2[i] = i + 5
	}
	intMeetInBoth := meetInBoth(arr, arr2)
	fmt.Println("3) ", intMeetInBoth)
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

//3.Есть два больших массива чисел, надо найти, какие упоминаются в обоих
func meetInBoth(first []int, second []int) []int {
	var res []int
	var midNums = make(map[int]available)
	for _, v := range first {
		midNums[v] = available{true, false}
	}

	for _, v := range second {
		_, okay := midNums[v]
		if okay != false {
			if midNums[v].isSecond == false {
				midNums[v] = available{true, true}
				res = append(res, v)
			}
		}
	}

	return res
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

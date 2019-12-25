package main

import (
	"fmt"
	//"math/rand"
	"strings"
	"unicode"
)

type available struct {
	isFirst  bool
	isSecond bool
}

func main2() {
	//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
	text := "Go (часто также Golang) — компилируемый многопоточный язык программирования, разработанный внутри компании Google. Разработка Go началась в сентябре 2007 года, его непосредственным проектированием занимались Роберт Гризмер, Роб Пайк и Кен Томпсон, занимавшиеся до этого проектом разработки операционной системы Inferno. Официально язык был представлен в ноябре 2009 года."
	fmt.Println("1) 	", wordCount(text))
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
	res := make(map[int]int, 5)
	res[1] = 0
	res[2] = 1
	print(calcFib(10, res))
	//	or
	//print(fib(9))

}

//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
func wordCount(str string) map[string]int {
	m := make(map[string]int)
	//strings.
	arr := strings.FieldsFunc(str, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	for _, v := range arr {
		m[v]++
	}
	return m
}

//2. Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func intCount(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
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

//func fib(n int) int {
//	res := []int{0, 1}
//	return calcFib(5, res)
//}

func calcFib(n int, res map[int]int) int {
	v, ok := res[n]
	if ok {
		return v
	}
	res[n] = calcFib(n-2, res) + calcFib(n-1, res)

	return res[n]
}

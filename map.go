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
	print(fib(10))
}

//1. Есть текст, надо посчитать сколько раз каждое слова встречается.
func wordCount(str string) map[string]int {
	arr := strings.FieldsFunc(str, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	// jekamas: объявлять переменные лучше ближе к месту их использования. тогда проще читать кусок кода.
	m := make(map[string]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

//2. Есть очень большой массив(слайс) целых чисел, надо сказать какие числа в нем упоминаются хоть по разу.
func intCount(arr []int) map[int]int {
	// jekamas: как вариант можно map[int]struct{} поскольку у нас не спрашивают подсчитать количество вхождений
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

//3.Есть два больших массива чисел, надо найти, какие упоминаются в обоих
func meetInBoth(first []int, second []int) []int {
	var midNums = make(map[int]available)
	for _, v := range first {
		midNums[v] = available{true, false}
	}

	var res []int
	for _, v := range second {
		//jekamas: чуть более общепринятые названия и выражения сделал
		if _, ok := midNums[v]; ok {
			if !midNums[v].isSecond {
				midNums[v] = available{true, true}
				res = append(res, v)
			}
		}
	}

	return res
}

func fib(n int) int {
	return calcFib(n, map[int]int{0: 0, 1: 1})
}

func calcFib(n int, res map[int]int) int {
	v, ok := res[n]
	if ok {
		return v
	}
	res[n] = calcFib(n-2, res) + calcFib(n-1, res)

	return res[n]
}

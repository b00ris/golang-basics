package main

import (
	"fmt"
	"sort"
)

func main2() {
	// 1. К каждому элементу []int прибавить 1
	arr := []int{0, 1, 2, 3}
	addOne(arr)
	fmt.Println("1) ", arr)
	// 2. Добавить в конец слайса число 5
	arr = append(arr, 5)
	fmt.Println("2) ", arr)
	// 3. Добавить в начало слайса число 5
	arr = append([]int{5}, arr...)
	fmt.Println("3) ", arr)
	// 4. Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	var value int
	arr, value = pop(arr)
	fmt.Println("4) ", value, "->", arr)
	//5. Взять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалить
	arr, value = firstPop(arr)
	fmt.Println("5) ", value, "->", arr)
	//6. Взять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию
	arr, value = popI(arr, 1)
	fmt.Println("6) ", value, "->", arr)
	//7. Объединить два слайса и вернуть новый со всеми элементами первого и второго
	arr2 := []int{1, 5, 4, 8, 9}
	arr3 := mergeArr(arr, arr2)
	fmt.Println("7) ", arr3)
	//8. Из первого слайса удалить все числа, которые есть во втором
	fmt.Print("8) ", arr, "-", arr2)
	arr = deleteDuplicate(arr, arr2)
	fmt.Println(" = ", arr)
	//9. Сдвинуть все элементы слайса на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
	fmt.Print("9) ", arr3)
	leftShift(arr3, 1)
	fmt.Println(" -> ", arr3)
	//10. Тоже, но сдвиг на заданное пользователем i
	fmt.Print("10) ", arr3)
	leftShift(arr3, 2)
	fmt.Println(" -> ", arr3)
	//11. Тоже, что 9, но сдвиг вправо
	fmt.Print("11) ", arr3)
	rightShift(arr3, 1)
	fmt.Println(" -> ", arr3)
	//12. Тоже, что 9, но сдвиг вправо
	fmt.Print("12) ", arr3)
	rightShift(arr3, 2)
	fmt.Println(" -> ", arr3)
	//13. Вернуть пользователю копию переданного слайса
	newArr := copyArr(arr3)
	fmt.Println("13) ", newArr)
	//14. В слайсе поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
	fmt.Print("14) ", arr3)
	alternationValue(arr3)
	fmt.Println(" -> ", arr3)
	//15. Упорядочить слайс в порядке: прямом, обратном, лексикографическом.
	fmt.Print("15) Возрастающая: ", arr3)
	sort.Ints(arr3)
	fmt.Println(" -> ", arr3)
	arr3 = copyArr(newArr)
	fmt.Print("Убывающая: ", arr3)
	sort.Slice(arr3, func(i, j int) bool {
		return arr3[i] > arr3[j]
	})
	fmt.Println(" -> ", arr3)
	strs := []string{"Hello", "Alfred", "Abc", "Holla"}
	fmt.Print("Лексикографическая:", strs)
	sort.Strings(strs)
	fmt.Println(" ->", strs)

}

//1. К каждому элементу []int прибавить 1
func addOne(arr []int) {
	for i := range arr {
		arr[i] += 1
	}
}

// 4. Взять последнее число слайса, вернуть его пользователю, а из слайса этот элемент удалить
func pop(arr []int) ([]int, int) {
	value := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	return arr, value
}

//5. Взять первое число слайса, вернуть его пользователю, а из слайса этот элемент удалить
func firstPop(arr []int) ([]int, int) {
	value := arr[0]
	arr = arr[1:]
	return arr, value
}

//6. Взять i-е число слайса, вернуть его пользователю, а из слайса этот элемент удалить. Число i передает пользователь в функцию
func popI(arr []int, i int) ([]int, int) {
	value := arr[i]
	arr = append(arr[:i], arr[i+1:]...)
	return arr, value
}

//7. Объединить два слайса и вернуть новый со всеми элементами первого и второго
func mergeArr(left, right []int) (arr []int) {
	arr = append(left, right...)
	return arr
}

//8. Из первого слайса удалить все числа, которые есть во втором
func deleteDuplicate(arr []int, arr2 []int) []int {
	//var resArr = arr
	for i, arrV := range arr {
		for _, v := range arr2 {
			if arrV == v {
				if i == 0 {
					arr = arr[1:]
				} else if i == len(arr) {
					arr = arr[:len(arr)-1]
				} else {
					arr = append(arr[:i], arr[i+1:]...)
				}
				break
			}
		}
	}
	return arr
}

//9-10. Сдвинуть все элементы слайса на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним.
func leftShift(arr []int, shift int) {
	tmp := make([]int, shift)
	copy(tmp, arr[0:shift])
	for i := 0; i < len(arr)-shift; i++ {
		arr[i] = arr[i+shift]

	}

	var j = 0
	for i := len(arr) - shift; i < len(arr); i++ {
		arr[i] = tmp[j]
		j++
	}
}

//11-12. Тоже, что 9, но сдвиг вправо
func rightShift(arr []int, shift int) {
	tmp := make([]int, shift)
	copy(tmp, arr[len(arr)-shift:])
	for i := len(arr) - 1; i > shift-1; i-- {
		arr[i] = arr[i-shift]
	}

	//var j = 0
	for i := 0; i < shift; i++ {
		arr[i] = tmp[i]
	}
}

//13. Вернуть пользователю копию переданного слайса
func copyArr(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}

//14. В слайсе поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
func alternationValue(arr []int) {
	for i := 0; i < len(arr)-1; i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

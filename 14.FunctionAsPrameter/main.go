package main

import "fmt"

func FilterNumber(data []int, filter func(int) bool) []int {
	var newNumberFilterd = []int{}
	for _, v := range data {
		if ok := filter(v); ok {
			newNumberFilterd = append(newNumberFilterd, v)
		}
	}
	return newNumberFilterd
}

type FilterEven func(int) bool

func FilterNumber2(data []int, callbacks FilterEven) []int {
	evenNumbers := []int{}
	for _, v := range data {
		if ok := callbacks(v); ok {
			evenNumbers = append(evenNumbers, v)
		}
	}
	return evenNumbers
}

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var dataLessThan5 = FilterNumber(data, func(each int) bool {
		return each < 5
	})

	var dataEven = FilterNumber2(data, func(i int) bool {
		return i%2 == 0
	})

	fmt.Println(dataLessThan5)
	fmt.Println(dataEven)
}

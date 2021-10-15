package main

import (
	"fmt"
	"math"
)

func main() {
	scorce := []int{80, 90, 100}

	// basic closure function
	var getHello = func() string {
		return "Hello"
	}

	// closure function with return value
	var avg = func(numbers ...int) float64 {
		a := 0
		for _, v := range numbers {
			a = a + v
		}
		return float64(a / len(numbers))
	}

	// closure with IIFE method
	numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	var filterNum = func(i int) []int {
		newNumber := []int{}
		for _, v := range numbers {
			if v < i {
				continue
			} else {
				newNumber = append(newNumber, v)
			}
		}
		return newNumber
	}(3)

	// closure function with multiple return value
	var calculate = func(d float64) (float64, float64) {
		var area = math.Pi * math.Pow(d/2, 2)
		var circumference = math.Pi * d
		return area, circumference
	}

	fmt.Println(getHello())
	fmt.Println(avg(scorce...))
	fmt.Println(calculate(20))
	fmt.Println(filterNum)
}

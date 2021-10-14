package main

import "fmt"

func FizzBuzz(numbers ...int) {
	for _, number := range numbers {
		if number%3 == 0 && number%5 == 0 {
			fmt.Printf("[%d] Fizz Buzz \n", number)
		} else if number%3 == 0 {
			fmt.Printf("[%d] Fizz\n", number)
		} else if number%5 == 0 {
			fmt.Printf("[%d] Buzz\n", number)
		} else {
			fmt.Printf("[%d]\n", number)
		}
	}
}

func Average(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

func main() {
	numbers := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	FizzBuzz(numbers...)

	fmt.Println("Average is \t:", Average(numbers...))
}

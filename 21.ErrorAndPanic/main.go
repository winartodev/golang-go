package main

import (
	"errors"
	"fmt"
)

func GetAvg(numbers []int) (int, error) {
	var a int = 0
	if len(numbers) < 1 {
		return 0, errors.New("number not be nill")
	} else {
		for _, v := range numbers {
			a = a + v
		}
	}
	return a, nil
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	result, err := GetAvg(numbers)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Average :", result)
}

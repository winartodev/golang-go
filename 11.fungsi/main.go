package main

import (
	"fmt"
	"math"
)

func BasicFunction1() {
	fmt.Println("Basic Function 1")
}

func BasicFunction2(nama string) {
	fmt.Println("Halo ", nama)
}

func BasicFunctionWithReturn1(nama string) string {
	return nama
}

func BasicFunctionWithReturn2(a, b int) int {
	return a + b
}

func BasicFunctionWithReturn3(x int) []int {
	isEven := []int{}
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			isEven = append(isEven, i)
		}
	}

	return isEven
}

func BasicFunctionWithMultipleValue1(a, b int) (int, int) {
	return a, b
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

func main() {
	var nama string = "winarto"
	BasicFunction1()
	BasicFunction2(nama)

	nama = BasicFunctionWithReturn1("Budi")
	fmt.Println("Halo ", nama)

	x := BasicFunctionWithReturn2(1, 2)
	fmt.Println(x)

	fmt.Println(BasicFunctionWithReturn2(1, 2))

	fmt.Println(BasicFunctionWithReturn3(10))

	a, b := BasicFunctionWithMultipleValue1(4, 5)
	fmt.Println(a)
	fmt.Println(b)

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)
}

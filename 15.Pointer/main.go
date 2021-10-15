package main

import "fmt"

func Change(number *int, value int) {
	*number = value
}

func main() {
	var number1 int = 4
	var number2 *int = &number1
	fmt.Println("number1 (val) \t:", number1)
	fmt.Println("number1 (addr) \t:", &number1)

	fmt.Println("number2 (val) \t:", *number2)
	fmt.Println("number2 (addr) \t:", number2)

	number1 = 5
	fmt.Println("\nCHANGE val number1 TO 5")

	fmt.Println("number1 (val) \t:", number1)
	fmt.Println("number1 (addr) \t:", &number1)

	fmt.Println("number2 (val) \t:", *number2)
	fmt.Println("number2 (addr) \t:", number2)

	Change(&number1, 10)

	fmt.Println("\nCHANGE val number1 TO 10")

	fmt.Println("number1 (val) \t:", number1)
	fmt.Println("number1 (addr) \t:", &number1)

	fmt.Println("number2 (val) \t:", *number2)
	fmt.Println("number2 (addr) \t:", number2)
}

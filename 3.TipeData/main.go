package main

import "fmt"

func main() {
	// tipe data numerik non-desimal
	var number1 int = 10
	var number2 int32 = 20
	var number3 int64 = 30

	// tipe data numerik desimal
	var number4 float32 = 80.6
	var number5 float64 = 90.5

	// tipe data string
	var nama string = "winarto"

	// tipe data boolean
	var toogleOn bool = true

	fmt.Println("TIPE DATA NUMBER NON DESIMAL")
	fmt.Println("number1\t:", number1)
	fmt.Println("number2\t:", number2)
	fmt.Println("number3\t:", number3)

	fmt.Println("\nTIPE DATA NUMBER DESIMAL")
	fmt.Println("number4\t:", number4)
	fmt.Println("number5\t:", number5)

	fmt.Println("\nTIPE DATA STRING")
	fmt.Println("nama\t:", nama)

	fmt.Println("\nTIPE DATA BOOLEAN")
	fmt.Println("toogle on\t:", toogleOn)
}

package main

import "fmt"

func main() {
	// operator penjumlahan
	a := 6
	b := 3

	penjumlahan := a + b
	pengurangan := a - b
	pembagian := a / b
	perkalian := a * b

	x := ((a + b) * (a / b)) - b

	isOn := true
	isOff := false

	isOnANDisOff := isOn && isOff
	isOnORisOff := isOn || isOff
	reverseIsOn := !isOn

	fmt.Println("penjumlahan \t:", penjumlahan)
	fmt.Println("pengurangan \t:", pengurangan)
	fmt.Println("pembagian \t:", pembagian)
	fmt.Println("perkalian \t:", perkalian)
	fmt.Println("x \t\t:", x)

	fmt.Println("isOn AND isOff \t:", isOnANDisOff)
	fmt.Println("isOn OR isOff \t:", isOnORisOff)
	fmt.Println("Reverse isOn \t:", reverseIsOn)

}

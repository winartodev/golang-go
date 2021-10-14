package main

import (
	"fmt"
)

func main() {
	const pi float64 = 3.14

	// pi = 3  // can error because pi is constat variable

	var r float64 = 9

	luasLingkaran := pi * (r * r)

	fmt.Println("LUAS Lingkaran \t:", luasLingkaran)
}

package main

import "fmt"

func main() {
	nilai := [6]int{80, 90, 60, 75, 66, 88}

	keranjangBuah := [4]string{
		"Apple",
		"Mangga",
		"Pisang",
		"Jambu",
	}

	gorengan := make([]string, 2)
	gorengan[0] = "tahu bunting"
	gorengan[1] = "pisang goreng"

	fmt.Println("Len Nilai \t:", len(nilai))
	fmt.Println("Index ke-1\t:", nilai[1])

	for _, v := range keranjangBuah {
		fmt.Println(v)
	}

	fmt.Println(gorengan)
}

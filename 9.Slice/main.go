package main

import "fmt"

func main() {
	nilai := []int{80, 90, 60, 75, 66, 88}

	keranjangBuah := []string{
		"Apple",
		"Mangga",
		"Pisang",
		"Jambu",
	}

	fmt.Println("Len Nilai \t:", len(nilai))
	fmt.Println("Index ke-1\t:", nilai[1])

	fmt.Println(len(keranjangBuah))
	fmt.Println(cap(keranjangBuah))
	fmt.Println("Buah Apple \t:", keranjangBuah[0])
	fmt.Println("slice kerangjang buah 1-3 \t:", keranjangBuah[1:])
}

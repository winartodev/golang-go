package main

import "fmt"

func main() {
	// declarate variable with keyword var
	var nama string = "winarto"
	var npm int = 1811010008

	// declarate variable with :=
	jurusan := "Teknik Informatika"

	// declarate multiple variable
	buah1, buah2, buah3 := "mangga", "aple", "pisang"

	// variable kosong
	_ = "asdf"

	fmt.Println("Nama \t\t:", nama)
	fmt.Println("NPM \t\t:", npm)
	fmt.Println("Jurusan \t:", jurusan)
	fmt.Println("\n==============NAMA BUAH================")
	fmt.Println("Buah 1 \t:", buah1)
	fmt.Println("Buah 2 \t:", buah2)
	fmt.Println("Buah 3 \t:", buah3)
}

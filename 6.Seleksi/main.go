package main

import "fmt"

func main() {
	nilai := 75

	if nilai >= 80 {
		fmt.Println("A")
	} else if nilai >= 75 && nilai < 80 {
		fmt.Println("B")
	} else if nilai >= 60 && nilai < 75 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	pilihBuah := 1

	switch pilihBuah {
	case 1:
		fmt.Println("Aple")
	case 2:
		fmt.Println("Pisang")
	case 3:
		fmt.Println("Jambu")
	default:
		fmt.Println("unkown")
	}

	nilai = 50

	switch {
	case nilai >= 80:
		fmt.Println("A")
	case nilai >= 75 && nilai < 80:
		fmt.Println("B")
	case nilai >= 60 && nilai < 75:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}

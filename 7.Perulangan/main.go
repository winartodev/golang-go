package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i)
	}

	buah := []string{"aple", "mangga", "pisang", "jambu"}
	for k, v := range buah {
		fmt.Printf("\nbuah %d -> %s", k, v)
	}

	point := []int{10, 20, 30, 40, 50}
	for _, v := range point {
		if v == 30 {
			fmt.Println("Founded")
		} else {
			fmt.Println("not found")
		}
	}
}

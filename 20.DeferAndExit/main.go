package main

import (
	"fmt"
	"os"
)

func main() {
	// basic defer
	fmt.Println("Halo")
	defer fmt.Println("Nama Saya Winarto")

	// basic exit
	for i := 1; i < 10; i++ {
		if i < 5 {
			fmt.Println(i)
		} else {
			os.Exit(1)
		}
	}

}

package rectangle

import (
	"fmt"
)

func luas(p, l float64) {
	result := p * l
	fmt.Println("Luas Persegi Panjang :", result)
}

func Keliling(p, l float64) {
	result := 2 * (p + l)
	fmt.Println("Keliling Persegi Panjang :", result)
}

package main

import (
	"fmt"
	"math"
)

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	jariJari float64
}

func (p Persegi) Luas() {
	result := 4 * p.sisi
	fmt.Println("Luas Persegi \t\t:", result)
}

func (p Persegi) Keliling() {
	result := math.Pow(p.sisi, 2)
	fmt.Println("Keliling Persegi \t:", result)
}

// declarate method with pointer
func (l *Lingkaran) Luas() {
	result := 2 * math.Pi * l.jariJari
	fmt.Println("Luas Lingkaran \t\t:", result)
}

func (l *Lingkaran) Keliling() {
	result := math.Pi * (math.Pow(l.jariJari, 2))
	fmt.Println("Keliling Lingkaran \t:", result)
}

func main() {
	bangunDatar1 := Persegi{sisi: 5}
	bangunDatar1.Luas()
	bangunDatar1.Keliling()

	bangunDatar2 := Lingkaran{jariJari: 4}
	bangunDatar2.Luas()
	bangunDatar2.Keliling()

}

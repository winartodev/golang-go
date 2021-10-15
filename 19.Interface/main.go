package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	jariJari float64
}

type PersegiPanjang struct {
	panjang float64
	lebar   float64
}

type person struct {
	ID   int
	Nama string
}

func (l *Lingkaran) Luas() float64 {
	result := math.Pi * (math.Pow(l.jariJari, 2))
	return result
}

func (l *Lingkaran) Keliling() float64 {
	result := 2 * math.Pi * l.jariJari
	return result
}

func (p *PersegiPanjang) Luas() float64 {
	result := p.panjang * p.lebar
	return result
}

func (p *PersegiPanjang) Keliling() float64 {
	result := 2 * (p.panjang + p.lebar)
	return result
}

func main() {
	var bangundatar Hitung

	bangundatar = &Lingkaran{jariJari: 3}
	fmt.Println("Luas Lingkaran : ", bangundatar.Luas())
	fmt.Println("Keliling Lingkaran : ", bangundatar.Keliling())

	bangundatar = &PersegiPanjang{panjang: 4, lebar: 3}
	fmt.Println("Luas Persegi Panjang : ", bangundatar.Luas())
	fmt.Println("Keliling Persegi Panjang : ", bangundatar.Keliling())

	// implement empty interface

	var random interface{}
	random = 2
	fmt.Println("random interface ", random.(int)*2)

	random = []string{"a", "b", "c"}
	fmt.Println("random interface ", random.([]string))

	random = &person{ID: 1, Nama: "Joni"}
	fmt.Println("random interface ", random.(*person).Nama)

	var keranjangBuah = []map[string]interface{}{
		{"fruits": "Apple", "total": 3},
		{"fruits": "Manggo", "total": 4},
		{"fruits": "Pinaple", "total": 2},
	}
	for _, v := range keranjangBuah {
		fmt.Println(v["fruits"])
	}
}

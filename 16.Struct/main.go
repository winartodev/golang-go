package main

import "fmt"

type Mahasiswa struct {
	Npm  int
	Nama string
}

type Salary struct {
	Basic int
}

type Employee struct {
	ID     int
	Name   string
	Salary Salary
}

func main() {
	winarto := Mahasiswa{}
	winarto.Npm = 1811010008
	winarto.Nama = "Winarto"
	fmt.Println(winarto)

	dataMahasiswa := []Mahasiswa{
		{1811010001, "A"},
		{1811010002, "B"},
		{1811010003, "C"},
		{1811010004, "D"},
	}

	for k, v := range dataMahasiswa {
		k += 1
		fmt.Printf("\nData Mahasiswa ke-%d\n", k)
		fmt.Println("NPM \t:", v.Npm)
		fmt.Println("Nama \t:", v.Nama)
	}

	budi := Employee{}
	budi.ID = 1
	budi.Name = "Budi"
	budi.Salary.Basic = 10000000

	fmt.Println(budi)
	fmt.Println(budi.Salary.Basic)

	Bukalapak := []Employee{
		{
			ID:   1,
			Name: "Anto",
			Salary: Salary{
				Basic: 10000000,
			},
		},
		{
			ID:   2,
			Name: "Budi",
			Salary: Salary{
				Basic: 15000000,
			},
		},
		{
			ID:   3,
			Name: "Santi",
			Salary: Salary{
				Basic: 12500000,
			},
		},
	}

	for k, v := range Bukalapak {
		k += 1
		fmt.Printf("\nEmployee %d\n", k)
		fmt.Println("ID \t:", v.ID)
		fmt.Println("Name \t:", v.Name)
		fmt.Println("Salary \t:", v.Salary.Basic)
	}
}

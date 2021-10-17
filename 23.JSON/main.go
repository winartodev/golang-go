package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
}

type Mahasiswa struct {
	Npm      int    `json:"npm"`
	Person   Person `json:"person"`
	Jurusan  string `json:"jurusan"`
	Semester int    `json:"semester"`
}

func main() {
	// decode JSON
	var jsonString = `{"nama":"winarto", "umur":21}`
	var jsonData = []byte(jsonString)

	var data Person

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Nama \t:", data.Nama)
	fmt.Println("Umur \t:", data.Nama)

	// encode json
	dataMahasiswa := []Mahasiswa{
		{
			Npm: 1811010008,
			Person: Person{
				Nama: "Winarto",
				Umur: 21,
			},
			Jurusan:  "TI",
			Semester: 5,
		},
		{
			Npm: 1811010009,
			Person: Person{
				Nama: "Jodi",
				Umur: 22,
			},
			Jurusan:  "SI",
			Semester: 3,
		},
	}

	jsonData, err = json.Marshal(dataMahasiswa)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonString = string(jsonData)
	fmt.Println(jsonString)
}

package main

import "fmt"

func main() {
	// basic map
	keranjangBuah := map[string]string{
		"buah1": "aple", "buah2": "mangga", "buah3": "pisang",
	}
	for key, value := range keranjangBuah {
		fmt.Printf("%s --> %s\n", key, value)
	}

	value, isExist := keranjangBuah["buah1"]
	if isExist {
		fmt.Printf("\n%s", value)
	} else {
		fmt.Println("Not Found")
	}

	// combine slice and map
	dataMahasiswa := []map[string]string{
		{"npm": "1811010008", "nama": "Winarto", "jurusan": "teknik informatika"},
		{"npm": "1811010010", "nama": "Budi", "jurusan": "Manajemen"},
		{"npm": "1811010011", "nama": "Joko", "jurusan": "Akuntansi"},
	}
	for k, v := range dataMahasiswa {
		fmt.Printf("\ndata mahasiswa ke-%d\n", k)
		fmt.Printf("NPM \t\t:%s\n", v["npm"])
		fmt.Printf("Nama \t\t:%s\n", v["nama"])
		fmt.Printf("Jurusan \t:%s\n", v["jurusan"])
	}
}

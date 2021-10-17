package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Mahasiswa struct {
	Npm      int    `json:"npm"`
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Semester int    `json:"semester"`
}

var BASE_URL = "http://localhost:8080"

func RequestGetAllMahasiswa() []*Mahasiswa {
	var err error
	var client = &http.Client{}
	var data []*Mahasiswa
	request, err := http.NewRequest("GET", BASE_URL+"/Mahasiswas", nil)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	return data
}

func RequestGetMahasiswaByNPM(npm string) Mahasiswa {
	var err error
	var client = &http.Client{}
	var data Mahasiswa

	request, err := http.NewRequest("GET", BASE_URL+"/Mahasiswa?npm="+npm, nil)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return data
	}
	return data
}

func main() {
	var mahasiswas = RequestGetAllMahasiswa()
	for _, mahasiswa := range mahasiswas {
		fmt.Printf("NPM: %d\t Nama: %s\t Jurusan: %s\n", mahasiswa.Npm, mahasiswa.Nama, mahasiswa.Jurusan)
	}

	var mahasiswa = RequestGetMahasiswaByNPM("1811010002")
	fmt.Println(mahasiswa)
}

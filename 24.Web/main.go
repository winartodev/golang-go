package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Mahasiswa struct {
	Npm      int    `json:"npm"`
	Nama     string `json:"nama"`
	Jurusan  string `json:"jurusan"`
	Semester int    `json:"semester"`
}

var DataMahasiswa = []Mahasiswa{
	{Npm: 1811010001, Nama: "Andi", Jurusan: "TI", Semester: 2},
	{Npm: 1811010002, Nama: "Budi", Jurusan: "SI", Semester: 1},
	{Npm: 1811010003, Nama: "Dosi", Jurusan: "AK", Semester: 3},
	{Npm: 1811010004, Nama: "Edi", Jurusan: "MI", Semester: 5},
}

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		result, err := json.Marshal(DataMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(result))
		return
	}
	http.Error(w, " ", http.StatusBadRequest)
}

func GetMahasiswaByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		npm, _ := strconv.Atoi(r.FormValue("npm"))
		for _, v := range DataMahasiswa {
			if v.Npm == npm {
				result, err := json.Marshal(v)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprint(w, string(result))
				return
			}
		}
		fmt.Fprintf(w, "npm %d not found", npm)
		http.Error(w, " ", http.StatusBadRequest)
	} else {
		AddMahasiswa(w, r)
	}
}

func AddMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		newData := json.NewDecoder(r.Body)
		var mahasiswa Mahasiswa
		err := newData.Decode(&mahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		DataMahasiswa = append(DataMahasiswa, mahasiswa)
		result, _ := json.Marshal(&mahasiswa)

		fmt.Fprint(w, string(result))
		return
	}
	http.Error(w, " ", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/Mahasiswas", GetAllMahasiswa)
	http.HandleFunc("/Mahasiswa", GetMahasiswaByID)
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

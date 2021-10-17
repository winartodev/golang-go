package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Mahasiswa struct {
	npm     int
	nama    string
	jurusan string
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_mahasiswa")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return db
}

func GetAllMahasiswa() {
	var db = Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_mahasiswa")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close()

	var result []Mahasiswa

	for rows.Next() {
		var each = Mahasiswa{}

		var err = rows.Scan(&each.npm, &each.nama, &each.jurusan)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for k, each := range result {
		k += 1
		fmt.Println("MAHASISWA ", k)
		fmt.Println("NPM : ", each.npm)
		fmt.Println("Nama : ", each.nama)
		fmt.Println("Jurusan : ", each.jurusan)
	}
}

func GetMahasiswaByNPM(npm int) {
	var db = Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_mahasiswa WHERE npm = ?", npm)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	defer rows.Close()

	var mahasiswa = []Mahasiswa{}

	for rows.Next() {
		var mhs Mahasiswa

		err := rows.Scan(&mhs.npm, &mhs.nama, &mhs.jurusan)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}

		mahasiswa = append(mahasiswa, mhs)
	}

	for _, v := range mahasiswa {
		fmt.Println("NPM :", v.npm)
		fmt.Println("Nama :", v.nama)
		fmt.Println("Jurusan :", v.jurusan)
	}
}

func AddMahasiswa(mhs *Mahasiswa) {
	var db = Connect()
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO tbl_mahasiswa (npm, nama, jurusan) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	result, err := statement.Exec(&mhs.npm, &mhs.nama, &mhs.jurusan)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	if rowAffected, _ := result.RowsAffected(); rowAffected == 1 {
		fmt.Printf("\nMahasiswa NPM %d Berhasil di Tambah\n", mhs.npm)
		fmt.Println("NPM :", mhs.npm)
		fmt.Println("Nama :", mhs.nama)
		fmt.Println("Jurusan :", mhs.jurusan)
	}
}

func main() {
	fmt.Println("=== GET ALL MAHASISWA ===")
	GetAllMahasiswa()

	fmt.Println("\n=== GET MAHASISWA BY NPM ===")
	GetMahasiswaByNPM(1811010008)

	fmt.Println("\n=== ADD MAHASISWA ===")
	var mhs = Mahasiswa{npm: 1811010010, nama: "BUDI", jurusan: "SI"}
	AddMahasiswa(&mhs)
}

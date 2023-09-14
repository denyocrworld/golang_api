package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Umur    int
	Jurusan string
}

func main() {
	// Inisialisasi struct sesuai dengan urutan field.
	mhs := Mahasiswa{
		Nama:    "Alice",
		Umur:    20,
		Jurusan: "Informatika",
	}
	fmt.Printf("Nama: %s, Umur: %d, Jurusan: %s\n", mhs.Nama, mhs.Umur, mhs.Jurusan)
}

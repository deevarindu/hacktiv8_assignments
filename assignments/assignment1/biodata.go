package main

import "fmt"

type Biodatas struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (b *Biodatas) printBiodata() {
	fmt.Printf("Absen\t\t: %d\nNama\t\t: %v\nAlamat\t\t: %v\nPekerjaan\t: %v\nAlasan\t\t: %v\n", b.Absen, b.Nama, b.Alamat, b.Pekerjaan, b.Alasan)
}

func getBiodata(biodatas []Biodatas, absen int) {
	for _, b := range biodatas {
		if b.Absen == absen {
			b.printBiodata()
			return
		}
	}
	fmt.Println("absen tidak ditemukan")
}

func main() {
	biodatas := []Biodatas{
		{
			Absen:     1,
			Nama:      "Mochamad Suhri Ainur Rifky",
			Alamat:    "Universitas Pembangunan Nasional Veteran Jawa Timur",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar",
		},
		{
			Absen:     2,
			Nama:      "MUHAMAD RESTU FADILLAH",
			Alamat:    "Universitas Komputer Indonesia",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     3,
			Nama:      "ALDI AGENG TRI SEFTIAN",
			Alamat:    "Universitas Pasundan",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     4,
			Nama:      "ARYO PANDAPOTAN SITOMPUL",
			Alamat:    "Universitas Terbuka",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     5,
			Nama:      "WILDAN SHAFA DIANDRA",
			Alamat:    "Institut Teknologi Sepuluh Nopember",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     6,
			Nama:      "M. AZRI RIYANDI",
			Alamat:    "Universitas Nusa Putra",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     7,
			Nama:      "WAHYU SETIAWAN",
			Alamat:    "Universitas Pendidikan Indonesia",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     8,
			Nama:      "HEZKYA NATANAEL RAMLI",
			Alamat:    "Universitas Mikroskil",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     9,
			Nama:      "MUHAMMAD FEBRI ANDANI",
			Alamat:    "Universitas Pamulang",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
		{
			Absen:     10,
			Nama:      "Juan Simon Damanik",
			Alamat:    "Universitas Dinamika Bangsa",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar juga",
		},
	}

	absen := 0
	fmt.Print("Masukkan absen : ")
	fmt.Scan(&absen)

	getBiodata(biodatas, absen)
}

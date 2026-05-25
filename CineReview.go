package main

import (
	"fmt"
)

type Film struct {
	Judul     string
	Genre     string
	Tahun     int
	Deskripsi string
	Rating    float64
}

func main() {

	var films []Film = []Film{
		{"The Conjuring", "Horror", 2013, "Teror_Rumah_Berhantu", 8.6},
		{"Project Hail Mary", "Sci-Fi", 2026, "Misi_Menyelamatkan_Bumi", 9.1},
		{"Arcane", "Animation", 2021, "Konflik_Piltover_dan_Zaun", 9.4},
		{"Kimi No Nawa", "Romance", 2016, "Pertukaran_Tubuh_Misterius", 8.9},
		{"Spider-Man 3", "Action", 2007, "Pertarungan_Melawan_Venom", 8.0},
	}

	var pilihan int

	for {

		fmt.Println("\n====| CineReview |====")
		fmt.Println("1. Tambah Film")
		fmt.Println("2. Edit Film")
		fmt.Println("3. Hapus Film")
		fmt.Println("4. Cari Film")
		fmt.Println("5. Urutkan Film")
		fmt.Println("6. Tampilkan Semua Film")
		fmt.Println("7. Statistik Film")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			//tambahFilm

		} else if pilihan == 2 {

			//editFilm

		} else if pilihan == 3 {

			//hapusFilm

		} else if pilihan == 4 {

			//cariFilm

		} else if pilihan == 5 {

			//urutFilm

		} else if pilihan == 6 {

			tampilFilm(films)

		} else if pilihan == 7 {

			//statistikFilm

		} else if pilihan == 8 {

			fmt.Println("Program selesai")
			break

		} else {

			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tampilFilm(films []Film) {

	var i int

	for i = 0; i < len(films); i++ {

		fmt.Println("----------------------")
		fmt.Println("No      :", i+1)
		fmt.Println("Judul   :", films[i].Judul)
		fmt.Println("Genre   :", films[i].Genre)
		fmt.Println("Tahun   :", films[i].Tahun)
		fmt.Println("Deskripsi :", films[i].Deskripsi)
		fmt.Println("Rating  :", films[i].Rating)
	}
}
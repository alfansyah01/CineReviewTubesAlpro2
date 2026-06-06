package main

import (
	"fmt"
	"strings"
)

var start bool = true

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

		if start {
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
			start = false
		} else {
			fmt.Print("----------------------")
			fmt.Print("\nKetik 0 untuk kembali ke menu utama : ")
			fmt.Scan(&pilihan)
		}

		if pilihan == 1 {

			tambahFilm(&films)

		} else if pilihan == 2 {

			editFilm(&films)

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

		} else if pilihan == 0 {
			start = true

		} else {

			fmt.Println("Pilihan", pilihan, "tidak tersedia")
		}
	}
}

func tampilFilm(films []Film) {

	var i int

	for i = 0; i < len(films); i++ {

		fmt.Println("----------------------")
		fmt.Println("No      	:", i+1)
		fmt.Println("Judul   	:", films[i].Judul)
		fmt.Println("Genre   	:", films[i].Genre)
		fmt.Println("Tahun   	:", films[i].Tahun)
		fmt.Println("Deskripsi  :", films[i].Deskripsi)
		fmt.Println("Rating  	:", films[i].Rating)
	}
}

func tambahFilm(films *[]Film) {

	var filmBaru Film
	var jumlah int

	fmt.Print("Input jumlah film yang ingin ditambahkan atau ketik '0' untuk kembali ke menu: ")
	fmt.Scan(&jumlah)

	if jumlah > 0 {
		for i := 1; jumlah >= i; i++ {
			fmt.Print("Masukkan Judul Film ke-", i, ": ")
			fmt.Scan(&filmBaru.Judul)

			fmt.Print("Masukkan Genre Film ke-", i, ": ")
			fmt.Scan(&filmBaru.Genre)

			fmt.Print("Masukkan Tahun Film ke-", i, ": ")
			fmt.Scan(&filmBaru.Tahun)

			fmt.Print("Masukkan Deskripsi Film ke-", i, ": ")
			fmt.Scan(&filmBaru.Deskripsi)

			fmt.Print("Masukkan Rating Film ke-", i, ": ")
			fmt.Scan(&filmBaru.Rating)

			*films = append(*films, filmBaru)
			fmt.Printf("Berhasil menambahkan film ke-%d!\n", i)
		}

		fmt.Println(jumlah, "Film berhasil ditambahkan")
	} else {
		start = true
	}
}

func editFilm(films *[]Film) {

	var nomor int
	var pilihan string
	var ulangi bool = true

	tampilFilm(*films)

	for ulangi {
		fmt.Println("----------------------")
		fmt.Print("Pilih nomor film yang ingin diedit atau ketik '0' untuk kembali ke menu : ")
		fmt.Scan(&nomor)

		nomor--

		if nomor == -1 {
			start = true
		} else if nomor >= 0 && nomor < len(*films) {

			fmt.Println("Ketik 1 untuk edit seluruh isi atau ketik kategori yang ingin di edit (Judul, Genre dll)")
			fmt.Scan(&pilihan)
			if pilihan == "1" {
				fmt.Print("Judul Baru  : ")
				fmt.Scan(&(*films)[nomor].Judul)

				fmt.Print("Genre Baru  : ")
				fmt.Scan(&(*films)[nomor].Genre)

				fmt.Print("Tahun Baru  : ")
				fmt.Scan(&(*films)[nomor].Tahun)

				fmt.Print("Deskripsi Baru : ")
				fmt.Scan(&(*films)[nomor].Deskripsi)

				fmt.Print("Rating Baru : ")
				fmt.Scan(&(*films)[nomor].Rating)
				fmt.Println("Data film berhasil diubah")
				fmt.Print("Apakah anda ingin edit data film lagi? (y/n): ")
				fmt.Scan(&pilihan)
					if pilihan == "n" {
						ulangi = false
						start = true
					}
			} else {
				switch strings.ToLower(pilihan) {
				case "judul":
					fmt.Print("Judul Baru  : ")
					fmt.Scan(&(*films)[nomor].Judul)
				case "genre":
					fmt.Print("Genre Baru  : ")
					fmt.Scan(&(*films)[nomor].Genre)
				case "tahun":
					fmt.Print("Tahun Baru  : ")
					fmt.Scan(&(*films)[nomor].Tahun)
				case "deskripsi":
					fmt.Print("Deskripsi Baru : ")
					fmt.Scan(&(*films)[nomor].Deskripsi)
				case "rating":
					fmt.Print("Rating Baru : ")
					fmt.Scan(&(*films)[nomor].Rating)
				default:
					fmt.Println("Pilihan tidak valid")
					return
				}
				fmt.Println("Data film berhasil diubah")
				fmt.Print("Apakah anda ingin edit data film lagi? (y/n): ")
				fmt.Scan(&pilihan)
					if pilihan == "n" {
						ulangi = false
						start = true
					}
			}

		} else {

			fmt.Println("Nomor film tidak valid")
		}
	}
}

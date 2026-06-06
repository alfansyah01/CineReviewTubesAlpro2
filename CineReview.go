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

			hapusFilm(&films)

		} else if pilihan == 4 {

			var keyword string

			fmt.Print("Cari Berdasarkan judul/genre atau ketik '0' untuk kembali ke menu : ")
			fmt.Scan(&keyword)

			cariFilm(films, keyword)

		} else if pilihan == 5 {
			

		var keyword string

	fmt.Print("\nUrutkan berdasarkan rating/tahun : ")
	fmt.Scan(&keyword)

	keyword = strings.ToLower(keyword)

	urutFilm(&films, keyword)

	tampilFilm(films)

		} else if pilihan == 6 {

			tampilFilm(films)

		} else if pilihan == 7 {

			statistikFilm(films)

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

func hapusFilm(films *[]Film) {

	var nomor int

	tampilFilm(*films)

	fmt.Print("Pilih nomor film yang ingin dihapus : ")
	fmt.Scan(&nomor)

	nomor--

	if nomor >= 0 && nomor < len(*films) {

		*films = append((*films)[:nomor], (*films)[nomor+1:]...)

		fmt.Println("Film berhasil dihapus")

	} else {

		fmt.Println("Nomor film tidak valid")
	}
}

func cariFilm(films []Film, keyword string) {

	var cari string
	var i int
	var found bool = false

	if strings.EqualFold(keyword, "judul") {

		fmt.Print("Masukkan judul film : ")
		fmt.Scan(&cari)

		for i = 0; i < len(films); i++ {

			if strings.EqualFold(films[i].Judul, cari) {

				fmt.Println("----------------------")
				fmt.Println("Judul  :", films[i].Judul)
				fmt.Println("Genre  :", films[i].Genre)
				fmt.Println("Tahun  :", films[i].Tahun)
				fmt.Println("Rating :", films[i].Rating)

				found = true
			}
		}

	} else if strings.EqualFold(keyword, "genre") {

		fmt.Print("Masukkan genre film : ")
		fmt.Scan(&cari)

		for i = 0; i < len(films); i++ {

			if strings.EqualFold(films[i].Genre, cari) {

				fmt.Println("----------------------")
				fmt.Println("Judul  :", films[i].Judul)
				fmt.Println("Genre  :", films[i].Genre)
				fmt.Println("Tahun  :", films[i].Tahun)
				fmt.Println("Rating :", films[i].Rating)

				found = true
			}
		}

	} else if keyword == "0"{
		start = true
	}else{
		fmt.Println("Pilihan tidak tersedia")
		return
	}

	if found == false {
		fmt.Println("Film tidak tersedia")
	}
}

func urutFilm(films *[]Film, keyword string){
	var i int
	var j int
	var max int
	var temp Film
	var key Film

	if  keyword == "rating"{
		for i = 0; i < len(*films)-1; i++ {

		max = i

		for j = i + 1; j < len(*films); j++ {

			if (*films)[j].Rating > (*films)[max].Rating {
				max = j
			}
		}

		temp = (*films)[i]
		(*films)[i] = (*films)[max]
		(*films)[max] = temp
		}
	}else if keyword == "tahun"{
		for i = 1; i < len(*films); i++ {

		key = (*films)[i]
		j = i - 1

		for j >= 0 && (*films)[j].Tahun > key.Tahun {

			(*films)[j+1] = (*films)[j]
			j--
		}

		(*films)[j+1] = key
		}
	}else{
		fmt.Print("Pilihan tidak tersedia")
	}
}

func statistikFilm(films []Film) {

	var i int
	var action, thriller, drama, scifi, lainnya int
	var totalRating float64

	if len(films) == 0 {

		fmt.Println("Belum ada data film")
		return
	}

	for i = 0; i < len(films); i++ {

		totalRating += films[i].Rating

		if strings.EqualFold(films[i].Genre, "Action") {

			action++

		} else if strings.EqualFold(films[i].Genre, "Thriller") {

			thriller++

		} else if strings.EqualFold(films[i].Genre, "Drama") {

			drama++

		} else if strings.EqualFold(films[i].Genre, "Sci-Fi") {

			scifi++

		} else {

			lainnya++
		}
	}

	fmt.Println("\n===== STATISTIK FILM =====")
	fmt.Println("Jumlah Film Action   :", action)
	fmt.Println("Jumlah Film Thriller :", thriller)
	fmt.Println("Jumlah Film Drama    :", drama)
	fmt.Println("Jumlah Film Sci-Fi   :", scifi)
	fmt.Println("Genre Lainnya        :", lainnya)

	fmt.Printf("Rata-rata Rating : %.2f\n",
		totalRating/float64(len(films)))
}

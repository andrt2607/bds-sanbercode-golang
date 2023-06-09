package main

import (
	"fmt"
)

func main() {
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini

	tambahDataFilm := func(nama, durasi, genre, tahun string) {
		film := map[string]string{
			"nama":   nama,
			"durasi": durasi,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

func luasPersegiPanjang(panjang int, lebar int) (result int) {
	result = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (result int) {
	result = 2 * (panjang + lebar)
	return
}

func volumeBalok(panjang int, lebar int, tinggi int) (result int) {
	result = panjang * lebar * tinggi
	return
}

func introduce(name string, gender string, job string, age string) (result string) {
	result = "Pak " + name + " adalah seorang " + job + " yang berusia " + age
	return
}

func buahFavorit(nama string, buah ...string) (result string) {
	result = "halo nama saya " + nama + " dan buah favorit saya adalah "
	for _, element := range buah {
		result += element + ", "
	}
	return
}

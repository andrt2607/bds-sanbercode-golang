package main

import (
	"fmt"
)

func main() {
	//soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	updateValueLuas(&luasLigkaran, 7)
	updateValueKeliling(&kelilingLingkaran, 7)

	fmt.Printf("ini luas %.2f ini keliling %.2f", luasLigkaran, kelilingLingkaran)

	//soal 2

	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//soal 3
	var buah = []string{}
	addBuah(&buah, "Jeruk")
	addBuah(&buah, "Semangka")
	addBuah(&buah, "Mangga")
	addBuah(&buah, "Strawberry")
	addBuah(&buah, "Durian")
	addBuah(&buah, "Manggis")
	addBuah(&buah, "Alpukat")
	for i, item := range buah {
		fmt.Printf("%d. %s \n", i+1, item)
	}

	//soal 4

	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		fmt.Printf("\n\n%d. title: %s\n duration: %s \n genre: %s \n year: %s", i+1, item["title"], item["duration"], item["genre"], item["year"])
	}
}

func updateValueLuas(luasLigkaran *float64, jariJari float64) {
	*luasLigkaran = 3.14 * jariJari * jariJari
}
func updateValueKeliling(kelilingLingkaran *float64, jariJari float64) {
	*kelilingLingkaran = 2 * 3.14 * jariJari
}

func introduce(sentence *string, name string, gender string, job string, age string) {
	pakbu := ""
	if gender == "laki-laki" {
		pakbu = "pak"
	} else {
		pakbu = "bu"
	}
	*sentence = pakbu + name + " adalah  seorang " + job + " yang berusia " + age + " tahun"
}

func addBuah(buah *[]string, item string) {
	*buah = append(*buah, item)
}

func tambahDataFilm(title string, duration string, genre string, releaseDate string, dataFilm *[]map[string]string) {
	newFilm := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     releaseDate,
	}
	*dataFilm = append(*dataFilm, newFilm)
}

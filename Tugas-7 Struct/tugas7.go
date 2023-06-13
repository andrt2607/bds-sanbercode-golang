// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	//soal1
	type buah struct {
		name       string
		color      string
		adaBijinya bool
		harga      int
	}

	var nanas = buah{"nanas", "kuning", false, 9000}
	var jeruk = buah{"jeruk", "oranye", true, 8000}
	var semangka = buah{"semangka", "hijau & merah", true, 10000}
	var pisang = buah{"pisang", "kuning", false, 5000}
	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	//soal 2

	var segitiga1 = segitiga{4, 2}
	fmt.Println(segitiga1.luasSegitiga())
	var persegi1 = persegi{2}
	fmt.Println(persegi1.luasPersegi())
	var persegiPanjang1 = persegiPanjang{4, 2}
	fmt.Println(persegiPanjang1.luasPersegiPanjang())

	//soal 3
	phone1 := phone{
		name:   "poco m4",
		brand:  "xiaomi",
		year:   2023,
		colors: []string{},
	}
	new := "biru"
	newLagi := "merah"
	phone1.addColor(&phone1, new)
	phone1.addColor(&phone1, newLagi)

	fmt.Println(phone1.colors)

	//soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	fmt.Println(dataFilm)

}

type movie struct {
	title    string
	duration int
	genre    string
	year     int
}

func tambahDataFilm(title string, duration int, genre string, year int, target *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*target = append(*target, newMovie)
}

type segitiga struct {
	alas   int
	tinggi int
}

func (s segitiga) luasSegitiga() (result float64) {
	result = 0.5 * float64(s.alas) * float64(s.tinggi)
	return
	// return
}

type persegi struct {
	sisi int
}

func (p persegi) luasPersegi() (result float64) {
	result = float64(p.sisi) * float64(p.sisi)
	return
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luasPersegiPanjang() (result float64) {
	result = float64(pp.panjang) * float64(pp.lebar)
	return
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p phone) addColor(target *phone, newColor string) {
	target.colors = append(target.colors, newColor)
}

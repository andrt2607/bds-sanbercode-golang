// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var tes1 hitungBangunDatar

	tes1 = segitigaSamaSisi{alas: 2, tinggi: 4}
	fmt.Println("luas ss : ", tes1.luas())
	fmt.Println("keliling ss : ", tes1.keliling())

	tes1 = persegiPanjang{panjang: 2, lebar: 4}
	fmt.Println("luas pp : ", tes1.luas())
	fmt.Println("keliling pp : ", tes1.keliling())

	var tes2 hitungBangunRuang

	tes2 = tabung{jariJari: 7, tinggi: 5}
	fmt.Printf("\nvolume tb: %.2f\n", tes2.volume())
	fmt.Println("luas permukaan tb: ", tes2.luasPermukaan())

	tes2 = balok{panjang: 7, lebar: 4, tinggi: 5}
	fmt.Println("volume bl: ", tes2.volume())
	fmt.Println("luas permukaan bl: ", tes2.luasPermukaan())

	//soal 2
	var tes3 myInterface

	tes3 = phone{
		name:   "pocom4",
		brand:  "xiaomi",
		year:   2023,
		colors: []string{},
	}

	tes3.printDetail()

	//soal 3

	fmt.Println("ini batas")
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	convertPertama := kumpulanAngkaPertama.([]int)
	convertKedua := kumpulanAngkaKedua.([]int)

	var resultPlus = kumpulanAngkaPertama.([]int)[0] + kumpulanAngkaPertama.([]int)[1] + kumpulanAngkaKedua.([]int)[0] + kumpulanAngkaKedua.([]int)[1]
	fmt.Printf("%s %d %d = %d", prefix, convertPertama, convertKedua, resultPlus)
	// tulis jawaban anda disini

}

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (ss segitigaSamaSisi) luas() int {
	return (ss.alas * ss.tinggi) / 2
}

func (ss segitigaSamaSisi) keliling() int {
	return (2 * ss.luas()) / ss.tinggi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2*pp.panjang + 2*pp.lebar
}

func (tb tabung) volume() float64 {
	return 3.14 * tb.jariJari * tb.jariJari * tb.tinggi
}

func (tb tabung) luasPermukaan() float64 {
	return 2 * 3.14 * tb.jariJari * (tb.jariJari + tb.tinggi)
}

func (bl balok) volume() float64 {
	return float64(bl.panjang * bl.lebar * bl.tinggi)
}

func (bl balok) luasPermukaan() float64 {
	return float64(2 * (bl.panjang*bl.lebar + bl.panjang*bl.tinggi + bl.lebar*bl.tinggi))
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type myInterface interface {
	printDetail()
}

func (p phone) printDetail() {
	fmt.Printf("name: %s\nbrand: %s\nyear: %d\ncolors: %s", p.name, p.brand, p.year, p.colors)
}

func luasPersegi(sisi int, boolVal bool) interface{} {

	luas := sisi * sisi

	if sisi != 0 && boolVal == true {
		return "luas persegi dengan sisi" + strconv.Itoa(sisi) + "cm adalah " + strconv.Itoa(luas) + " cm"
	} else if sisi == 0 && boolVal == true {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && boolVal == false {
		return nil
	}
	return sisi

}

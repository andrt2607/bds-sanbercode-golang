package main

import (
	"fmt"

	. "./hitung"
	// _ "./hitung"
)

func main() {
	//soal 1
	var tes1 HitungBangunDatar

	tes1 = SegitigaSamaSisi{Alas: 2, Tinggi: 4}
	fmt.Println("luas ss : ", tes1.Luas())
	fmt.Println("keliling ss : ", tes1.Keliling())

	tes1 = PersegiPanjang{Panjang: 2, Lebar: 4}
	fmt.Println("luas pp : ", tes1.Luas())
	fmt.Println("keliling pp : ", tes1.Keliling())

	var tes2 HitungBangunRuang

	tes2 = Tabung{JariJari: 7, Tinggi: 5}
	fmt.Printf("\nvolume tb: %.2f\n", tes2.Volume())
	fmt.Println("luas permukaan tb: ", tes2.LuasPermukaan())

	tes2 = Balok{Panjang: 7, Lebar: 4, Tinggi: 5}
	fmt.Println("volume bl: ", tes2.Volume())
	fmt.Println("luas permukaan bl: ", tes2.LuasPermukaan())

	//soal 2
	var tes3 MyInterface

	tes3 = Phone{
		Name:   "pocom4",
		Brand:  "xiaomi",
		Year:   2023,
		Colors: []string{},
	}

	tes3.PrintDetail()

	//soal 3

	fmt.Println("ini batas")
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	convertPertama := kumpulanAngkaPertama.([]int)
	convertKedua := kumpulanAngkaKedua.([]int)

	var resultPlus = kumpulanAngkaPertama.([]int)[0] + kumpulanAngkaPertama.([]int)[1] + kumpulanAngkaKedua.([]int)[0] + kumpulanAngkaKedua.([]int)[1]
	fmt.Printf("%s %d %d = %d", prefix, convertPertama, convertKedua, resultPlus)
}

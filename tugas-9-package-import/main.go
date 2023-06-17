package main

import (
	"fmt"

	ht "tugas9/Hitung"
	// _ "./hitung"
)

func main() {
	//soal 1
	var tes1 ht.HitungBangunDatar

	tes1 = ht.SegitigaSamaSisi{Alas: 2, Tinggi: 4}
	fmt.Println("luas ss : ", tes1.Luas())
	fmt.Println("keliling ss : ", tes1.Keliling())

	tes1 = ht.PersegiPanjang{Panjang: 2, Lebar: 4}
	fmt.Println("luas pp : ", tes1.Luas())
	fmt.Println("keliling pp : ", tes1.Keliling())

	var tes2 ht.HitungBangunRuang

	tes2 = ht.Tabung{JariJari: 7, Tinggi: 5}
	fmt.Printf("\nvolume tb: %.2f\n", tes2.Volume())
	fmt.Println("luas permukaan tb: ", tes2.LuasPermukaan())

	tes2 = ht.Balok{Panjang: 7, Lebar: 4, Tinggi: 5}
	fmt.Println("volume bl: ", tes2.Volume())
	fmt.Println("luas permukaan bl: ", tes2.LuasPermukaan())

	//soal 2
	var tes3 ht.MyInterface

	tes3 = ht.Phone{
		Name:   "pocom4",
		Brand:  "xiaomi",
		Year:   2023,
		Colors: []string{},
	}

	tes3.PrintDetail()

	//soal 3

	fmt.Println("ini batas")
	fmt.Println(ht.LuasPersegi(4, true))

	fmt.Println(ht.LuasPersegi(8, false))

	fmt.Println(ht.LuasPersegi(0, true))

	fmt.Println(ht.LuasPersegi(0, false))

	//soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	convertPertama := kumpulanAngkaPertama.([]int)
	convertKedua := kumpulanAngkaKedua.([]int)

	var resultPlus = kumpulanAngkaPertama.([]int)[0] + kumpulanAngkaPertama.([]int)[1] + kumpulanAngkaKedua.([]int)[0] + kumpulanAngkaKedua.([]int)[1]
	fmt.Printf("%s %d %d = %d", prefix, convertPertama, convertKedua, resultPlus)
}

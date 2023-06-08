package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var convertPanjang, _ = strconv.Atoi(panjangPersegiPanjang)
	var convertLebar, _ = strconv.Atoi(lebarPersegiPanjang)

	var luasPersegiPanjang int = convertPanjang * convertLebar

	var convertAlas, _ = strconv.Atoi(alasSegitiga)
	var convertTinggi, _ = strconv.Atoi(tinggiSegitiga)

	var kelilingPersegiPanjang int = 2 * (convertPanjang + convertLebar)
	var luasSegitiga int = 1 / 2 * convertAlas * convertTinggi

	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	//soal 2

	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("Ini nilaiJohn")

	if nilaiJohn >= 80 {
		fmt.Println("indeks A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeks B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeks C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeks D")
	} else {
		fmt.Println("indeks E")
	}

	fmt.Println("Ini nilaiDoe")
	if nilaiDoe >= 80 {
		fmt.Println("indeks A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("indeks B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("indeks C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("indeks D")
	} else {
		fmt.Println("indeks E")
	}

	//soal 3
	var tanggal = 27
	var bulan = 6
	var tahun = 1998

	var converTanggal = strconv.Itoa(tanggal)
	convertBulan := ""
	var convertTahun = strconv.Itoa(tahun)

	switch bulan {
	case 1:
		convertBulan = "Januari"
	case 2:
		convertBulan = "Februari"
	case 3:
		convertBulan = "Maret"
	case 4:
		convertBulan = "April"
	case 5:
		convertBulan = "Mei"
	case 6:
		convertBulan = "Juni"
	case 7:
		convertBulan = "Juli"
	case 8:
		convertBulan = "Agustus"
	case 9:
		convertBulan = "September"
	case 10:
		convertBulan = "Oktober"
	case 11:
		convertBulan = "November"
	case 12:
		convertBulan = "Desember"
	}

	fmt.Println(converTanggal + convertBulan + convertTahun)

	//soal 4
	switch {
	case tahun > 1944 && tahun < 1964:
		fmt.Println("Baby Boomer")
	case tahun > 1965 && tahun < 1979:
		fmt.Println("Generasi X")
	case tahun > 1980 && tahun < 1994:
		fmt.Println("Generasi Y")
	case tahun > 1995 && tahun < 2015:
		fmt.Println("Generasi Z")
	}
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//soal 1
	fmt.Println("Bootcamp Digital Skill Sanbercode Golang")

	//soal 2
	halo := "Halo Dunia"

	// var newConvert = strings.Replace(halo, "Dunia", "Golang", 5)
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", 5))

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(kataPertama + " " + strings.Replace(kataKedua, "s", "S", 1) + " " + strings.ToLower(kataKetiga) + " " + strings.ToUpper(kataKeempat))

	//soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var convPertama, _ = strconv.Atoi(angkaPertama)
	var convKedua, _ = strconv.Atoi(angkaKedua)
	var convKetiga, _ = strconv.Atoi(angkaKetiga)
	var convKeempat, _ = strconv.Atoi(angkaKeempat)
	fmt.Println(convPertama + convKedua + convKetiga + convKeempat)

	//soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var newText = strings.Replace(kalimat, "halo", "Hi", 2)
	resultText := newText + " - " + strconv.Itoa(angka)
	fmt.Println(resultText)

}

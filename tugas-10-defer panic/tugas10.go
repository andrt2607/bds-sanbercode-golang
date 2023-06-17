package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)

	// fmt.Println("tipe  variabel :", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variabel :", reflectValue.Int())
	// }
	//soal 2
	defer TestDefer("Golang Backend Development", 2021)

	fmt.Println(kelilingSegitigaSamaSisi(4, true))

	fmt.Println(kelilingSegitigaSamaSisi(8, false))

	fmt.Println(kelilingSegitigaSamaSisi(0, true))

	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// deklarasi variabel angka ini simpan di baris pertama func main

	//soal 3
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//soal 4

	var phones = []string{}

	tambahPhone("Xiaomi", &phones)
	tambahPhone("Asus", &phones)
	tambahPhone("Iphone", &phones)
	tambahPhone("Samsung", &phones)
	tambahPhone("Oppo", &phones)
	tambahPhone("Realme", &phones)
	tambahPhone("Vivo", &phones)

	luasLingkaran(7)
	luasLingkaran(10)
	luasLingkaran(15)
	kelilingLingkaran(7)
	kelilingLingkaran(10)
	kelilingLingkaran(15)

	// var name = flag.String("name", "anonymous", "type your name")
	// var age = flag.Int64("age", 25, "type your age")
	var panjang = flag.Int("panjang", 5, "type your panjang")
	var lebar = flag.Int("lebar", 4, "type your lebar")

	flag.Parse()
	fmt.Println(luasPersegiPanjang(*panjang, *lebar))
	fmt.Println(kelilingPersegiPanjang(*panjang, *lebar))

}

// soal 1
func TestDefer(kalimat string, tahun int) {
	message := recover()
	fmt.Println("Terjadi error ", message)
	fmt.Println(kalimat + strconv.Itoa(tahun))
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, boolValue bool) interface{} {
	result := 4 * 3
	if boolValue == true && sisi == 0 {
		// defer TestDefer("Golang Backend Development", 2021)
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	} else if boolValue == true {
		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(result) + " cm"
	} else if boolValue == false && sisi == 0 {
		defer TestDefer("Golang Backend Development", 2021)
		panic("Maaf anda belum menginput sisi dari segitiga sama sisi")

	}
	return strconv.Itoa(result)
}

// soal 3
func tambahAngka(number int, angka *int) {
	*angka += number
}

func cetakAngka(angka *int) {
	fmt.Println(strconv.Itoa(*angka))
}

func tambahPhone(object string, phones *[]string) {
	time.Sleep(time.Second * 1)
	*phones = append(*phones, object)
	fmt.Println(strconv.Itoa(len(*phones)) + ". " + object)
}

func luasLingkaran(jariJari float64) {
	result := math.Phi * jariJari
	fmt.Println(math.Round(result))
}

func kelilingLingkaran(jariJari float64) {
	result := 2 * math.Phi * jariJari
	fmt.Println(math.Round(result))
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2*panjang + 2*lebar
}

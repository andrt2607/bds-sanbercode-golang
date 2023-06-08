package main

import (
	"fmt"
	"strconv"
)

func main() {
	//soal 1
	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + " Berkualitas")
		} else if i%3 == 0 {
			fmt.Println(strconv.Itoa(i) + " I Love Coding")
		} else {
			fmt.Println(strconv.Itoa(i) + " Santai")
		}
	}

	//soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("#")
		}
		fmt.Println(" ")
	}

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:7])

	//soal 4
	var sayuran = []string{}
	sayuran = append(sayuran, "Bayam")
	sayuran = append(sayuran, "Buncis")
	sayuran = append(sayuran, "Kangkung")
	sayuran = append(sayuran, "Kubis")
	sayuran = append(sayuran, "Seledri")
	sayuran = append(sayuran, "Tauge")
	sayuran = append(sayuran, "Timun")

	for i, element := range sayuran {
		fmt.Println(strconv.Itoa(i+1) + " " + element)
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, element := range satuan {
		fmt.Println(key + " = " + strconv.Itoa(element))
	}
	fmt.Println("Volume Balok = " + strconv.Itoa((satuan["panjang"] * satuan["lebar"] * satuan["tinggi"])))

}

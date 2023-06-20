package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

func main() {
	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("numbers :", numbers)

	// var ch1 = make(chan float64)
	// go getAverage(numbers, ch1)

	// var ch2 = make(chan int)
	// go getMax(numbers, ch2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case avg := <-ch1:
	// 		fmt.Printf("Avg \t: %.2f \n", avg)
	// 	case max := <-ch2:
	// 		fmt.Printf("Max \t: %d \n", max)
	// 	}
	// }
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go printText("Halo", &wg)

	// wg.Add(1)
	// go printText("Dunia", &wg)

	// wg.Wait()

	//soal 1

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var myWg sync.WaitGroup
	myWg.Add(1)
	go printPhoneText(phones, &myWg)
	myWg.Wait()

	//soal 2

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	//soal 3

	luasLingkaran := make(chan float64, 3)

	go countLuasLingkaran(luasLingkaran, 8)
	go countLuasLingkaran(luasLingkaran, 14)
	go countLuasLingkaran(luasLingkaran, 20)

	// fmt.Println(<-luasLingkaran)
	// fmt.Println(<-luasLingkaran)
	// fmt.Println(<-luasLingkaran)

	kelilingLingkaran := make(chan float64, 3)
	go countKelilingLingkaran(kelilingLingkaran, 8)
	go countKelilingLingkaran(kelilingLingkaran, 14)
	go countKelilingLingkaran(kelilingLingkaran, 20)

	// fmt.Println(<-kelilingLingkaran)
	// fmt.Println(<-kelilingLingkaran)
	// fmt.Println(<-kelilingLingkaran)

	volumeTabung := make(chan float64, 3)
	go countVolumeTabung(volumeTabung, 8, 10)
	go countVolumeTabung(volumeTabung, 14, 10)
	go countVolumeTabung(volumeTabung, 20, 10)

	// fmt.Println(<-volumeTabung)
	// fmt.Println(<-volumeTabung)
	// fmt.Println(<-volumeTabung)

	for i := 0; i < 9; i++ {
		select {
		case keliling := <-kelilingLingkaran:
			fmt.Println("ini keliling ling: " + strconv.Itoa(int(keliling)))
		case luas := <-luasLingkaran:
			fmt.Println("ini luas ling: " + strconv.Itoa(int(luas)))
		case volume := <-volumeTabung:
			fmt.Println("ini volume ling: " + strconv.Itoa(int(volume)))
		}
	}

	luasPersegiPanjang := make(chan int)
	go countLuasPersegiPanjang(luasPersegiPanjang, 6, 4)
	kelilingPersegiPanjang := make(chan int)
	go countKelilingPersegiPanjang(kelilingPersegiPanjang, 6, 4)
	volumeBalok := make(chan int)
	go countVolumeBalok(volumeBalok, 6, 4, 10)

	for j := 0; j < 3; j++ {
		select {
		case luasPP := <-luasPersegiPanjang:
			fmt.Println("ini luas PP : " + strconv.Itoa(luasPP))
		case kelPP := <-kelilingPersegiPanjang:
			fmt.Println("ini keliling PP : " + strconv.Itoa(kelPP))
		case volPP := <-volumeBalok:
			fmt.Println("ini volum balok : " + strconv.Itoa(volPP))

		}
	}

}

func countLuasPersegiPanjang(ch chan<- int, panjang int, lebar int) {
	result := panjang * lebar
	ch <- result
}

func countKelilingPersegiPanjang(ch chan<- int, panjang int, lebar int) {
	result := 2*panjang + 2*lebar
	ch <- result
}

func countVolumeBalok(ch chan<- int, panjang int, lebar int, tinggi int) {
	result := panjang * lebar * tinggi
	ch <- result
}

func countVolumeTabung(ch chan<- float64, jariJari float64, tinggi float64) {
	result := math.Phi * jariJari * jariJari * tinggi
	ch <- math.Round(result)
}

func countKelilingLingkaran(ch chan<- float64, jariJari float64) {
	result := 2 * math.Phi * jariJari
	ch <- math.Round(result)
}

func countLuasLingkaran(ch chan<- float64, jariJari float64) {
	result := math.Phi * jariJari * jariJari
	ch <- math.Round(result)
	// close(ch)
}

func getMovies(ch chan<- string, movies ...string) {
	for i, item := range movies {
		ch <- strconv.Itoa(i+1) + ". " + item
	}
	close(ch)
}

func printPhoneText(phones []string, myWg *sync.WaitGroup) {
	for i, item := range phones {
		// myWg.Wait()
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(strconv.Itoa(i+1) + ". " + item)
	}
	myWg.Done()
}

// func getAverage(numbers []int, ch chan float64) {
// 	var sum = 0
// 	for _, e := range numbers {
// 		sum += e
// 	}
// 	ch <- float64(sum) / float64(len(numbers))
// }

// func getMax(numbers []int, ch chan int) {
// 	var max = numbers[0]
// 	for _, e := range numbers {
// 		if max < e {
// 			max = e
// 		}
// 	}
// 	ch <- max
// }

// func printText(text string, wg *sync.WaitGroup) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(text)
// 	}
// 	wg.Done()
// }

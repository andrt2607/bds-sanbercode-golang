package hitung

import (
	"fmt"
	"strconv"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (ss SegitigaSamaSisi) Luas() int {
	return (ss.Alas * ss.Tinggi) / 2
}

func (ss SegitigaSamaSisi) Keliling() int {
	return (2 * ss.Luas()) / ss.Tinggi
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return 2*pp.Panjang + 2*pp.Lebar
}

func (tb Tabung) Volume() float64 {
	return 3.14 * tb.JariJari * tb.JariJari * tb.Tinggi
}

func (tb Tabung) LuasPermukaan() float64 {
	return 2 * 3.14 * tb.JariJari * (tb.JariJari + tb.Tinggi)
}

func (bl Balok) Volume() float64 {
	return float64(bl.Panjang * bl.Lebar * bl.Tinggi)
}

func (bl Balok) LuasPermukaan() float64 {
	return float64(2 * (bl.Panjang*bl.Lebar + bl.Panjang*bl.Tinggi + bl.Lebar*bl.Tinggi))
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type MyInterface interface {
	PrintDetail()
}

func (p Phone) PrintDetail() {
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %s", p.Name, p.Brand, p.Year, p.Colors)
}

func LuasPersegi(sisi int, boolVal bool) interface{} {

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

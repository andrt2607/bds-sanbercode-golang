package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"

	config "quiz3/Config"
	"quiz3/models"
	"quiz3/repo"
	"quiz3/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}
	router := httprouter.New()
	router.GET("/segitiga-sama-sisi", CountSegitigaSamaSisi)
	router.GET("/persegi", CountPersegi)
	router.GET("/persegi-panjang", CountPersegiPanjang)
	router.GET("/lingkaran", CountLingkaran)
	router.GET("/jajar-genjang", CountJajarGenjang)
	router.GET("/categories", GetCategories)
	router.GET("/categories/:id/books", GetAllBooksByCategory)
	// router.Handler("GET","/categories", Auth(PostCategory))
	router.POST("/categories", Auth(PostCategory))
	router.PUT("/categories/:id", Auth(UpdateCategory))
	router.DELETE("/categories/:id", Auth(DeleteCategory))
	router.GET("/books", GetBooks)
	// router.handl
	// router.GET("/categories/:id/books", GetAllBooksByCategory)
	router.POST("/books", Auth(PostBook))
	router.PUT("/books/:id", Auth(UpdateBook))
	router.DELETE("/books/:id", Auth(DeleteBook))
	fmt.Println("---Server is running---")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			// w.Write([]byte("Username atau Password tidak boleh kosong"))
			// return
			defer ErrorHandler(w)
			panic("Username atau Password tidak boleh kosong")
		}

		if uname == "admin" && pwd == "password" {
			next(w, r, ps)
			return
		} else if uname == "editor" && pwd == "secret" {
			next(w, r, ps)
			return
		} else if uname == "trainer" && pwd == "rahasia" {
			next(w, r, ps)
			return
		}
		// w.Write([]byte("Username atau Password tidak sesuai"))
		defer ErrorHandler(w)
		panic("Username atau Password tidak sesuai")
	}
}

func CountSegitigaSamaSisi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mychan := make(chan float64)
	alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
	tinggi, _ := strconv.Atoi(r.URL.Query().Get("tinggi"))
	switch {
	case r.URL.Query().Get("hitung") == "luas":
		go operateLuasSS(mychan, alas, tinggi)
	case r.URL.Query().Get("hitung") == "keliling":
		go operateKelilingSS(mychan, alas)
	}

	select {
	case result := <-mychan:
		res := map[string]interface{}{
			"message": "Berhasil hitung segitiga sama sisi",
			"data":    result,
		}
		utils.ResponseJSON(w, res, http.StatusOK)
	}

}

func operateLuasSS(ch chan<- float64, alas int, tinggi int) {
	result := math.Round(float64(alas * tinggi * 2))
	ch <- result
}
func operateKelilingSS(ch chan<- float64, alas int) {
	result := math.Round(float64(alas * 3))
	ch <- result
}

func CountPersegi(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	other := make(chan float64)
	sisi, _ := strconv.Atoi(r.URL.Query().Get("sisi"))
	switch {
	case r.URL.Query().Get("hitung") == "luas":
		go operateLuasPersegi(other, sisi)
	case r.URL.Query().Get("hitung") == "keliling":
		go operateKelilingPersegi(other, sisi)
	}

	select {
	case result := <-other:
		res := map[string]interface{}{
			"message": "Berhasil hitung persegi",
			"data":    result,
		}
		utils.ResponseJSON(w, res, http.StatusOK)
	}

}

func operateLuasPersegi(ch chan<- float64, sisi int) {
	result := math.Round(float64(sisi * sisi))
	ch <- result
}
func operateKelilingPersegi(ch chan<- float64, sisi int) {
	result := math.Round(float64(sisi * 4))
	ch <- result
}

func CountPersegiPanjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mychan := make(chan float64)
	panjang, _ := strconv.Atoi(r.URL.Query().Get("panjang"))
	lebar, _ := strconv.Atoi(r.URL.Query().Get("lebar"))
	switch {
	case r.URL.Query().Get("hitung") == "luas":
		go operateLuasPersegiPanjang(mychan, panjang, lebar)
	case r.URL.Query().Get("hitung") == "keliling":
		go operateKelilingPersegiPanjang(mychan, panjang, lebar)
	}

	select {
	case result := <-mychan:
		res := map[string]interface{}{
			"message": "Berhasil hitung persegi panjang",
			"data":    result,
		}
		utils.ResponseJSON(w, res, http.StatusOK)
	}

}

func operateLuasPersegiPanjang(ch chan<- float64, panjang int, lebar int) {
	result := math.Round(float64(panjang * lebar))
	// fmt.Println("ini luas pp : ", panjang*lebar)
	ch <- result
}
func operateKelilingPersegiPanjang(ch chan<- float64, panjang int, lebar int) {
	result := math.Round(float64(2*panjang + 2*lebar))
	ch <- result
}

func CountLingkaran(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mychan := make(chan float64)
	jariJari, _ := strconv.Atoi(r.URL.Query().Get("jarijari"))
	switch {
	case r.URL.Query().Get("hitung") == "luas":
		go operateLuasLingkaran(mychan, jariJari)
	case r.URL.Query().Get("hitung") == "keliling":
		go operateKelilingLingkaran(mychan, jariJari)
	}

	select {
	case result := <-mychan:
		res := map[string]interface{}{
			"message": "Berhasil hitung lingkaran",
			"data":    result,
		}
		utils.ResponseJSON(w, res, http.StatusOK)
	}

}

func operateLuasLingkaran(ch chan<- float64, jariJari int) {
	result := math.Round(float64(math.Phi * float64(jariJari) * float64(jariJari)))
	fmt.Println("ini luas lingkaran : ", result)
	ch <- result
}
func operateKelilingLingkaran(ch chan<- float64, jariJari int) {
	result := math.Round(float64(2 * math.Phi * float64(jariJari)))
	ch <- result
}

func CountJajarGenjang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mychan := make(chan float64)
	sisi, _ := strconv.Atoi(r.URL.Query().Get("sisi"))
	alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
	tinggi, _ := strconv.Atoi(r.URL.Query().Get("tinggi"))
	switch {
	case r.URL.Query().Get("hitung") == "luas":
		go operateLuasJajarGenjang(mychan, alas, tinggi)
	case r.URL.Query().Get("hitung") == "keliling":
		go operateKelilingJajarGenjang(mychan, sisi, alas)
	}

	select {
	case result := <-mychan:
		res := map[string]interface{}{
			"message": "Berhasil hitung jajar genjang",
			"data":    result,
		}
		utils.ResponseJSON(w, res, http.StatusOK)
	}

}

func operateLuasJajarGenjang(ch chan<- float64, alas int, tinggi int) {
	result := math.Round(float64(alas * tinggi))
	ch <- result
}
func operateKelilingJajarGenjang(ch chan<- float64, alas int, sisi int) {
	result := math.Round(float64(2*alas + 2*sisi))
	ch <- result
}

func GetCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	categories, err := repo.GetAllCategories(ctx)

	if err != nil {
		fmt.Println(err)
	}
	res := map[string]interface{}{
		"message": "Data categories ditemukan",
		"data":    categories,
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var itemCategory models.Category

	if err := json.NewDecoder(r.Body).Decode(&itemCategory); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repo.Insertcategory(ctx, itemCategory); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"message": "berhasil tambah category",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	//untuk mengambil matkul params
	var idcategory, _ = strconv.Atoi(ps.ByName("id"))

	if err := repo.Updatecategory(ctx, category, idcategory); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"message": fmt.Sprintf("update category %d sukses", idcategory),
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory, _ = strconv.Atoi(ps.ByName("id"))
	if err := repo.Deletecategory(ctx, idCategory); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"message": fmt.Sprintf("delete category %d sukses", idCategory),
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func GetAllBooksByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// var idCategory, _ = strconv.Atoi(ps.ByName("id"))
	// if err := repo.Deletecategory(ctx, idCategory); err != nil {
	// 	kesalahan := map[string]string{
	// 		"error": fmt.Sprintf("%v", err),
	// 	}
	// 	utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
	// 	return
	// }
	// res := map[string]string{
	// 	"message": fmt.Sprintf("delete category %d sukses", idCategory),
	// }
	// utils.ResponseJSON(w, res, http.StatusOK)
}

func GetBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	books, err := repo.GetAllBooks(ctx)

	if err != nil {
		fmt.Println(err)
	}
	res := map[string]interface{}{
		"message": "Data books ditemukan",
		"data":    books,
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var itemBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&itemBook); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	// _, err := url.ParseRequestURI(itemBook.ImageUrl)
	// switch {
	// case err != nil:
	// 	// fallthrough
	// 	panic("image url harus berupa uri")
	// case itemBook.ReleaseYear < 1980 && itemBook.ReleaseYear > 2021:
	// 	panic("release year harus antara 1980 sampai 2021")
	// default:
	// 	fmt.Println("akan insert book")
	// 	if err := repo.InsertBook(ctx, itemBook); err != nil {
	// 		utils.ResponseJSON(w, err, http.StatusInternalServerError)
	// 		return
	// 	}
	// }
	if _, err := url.ParseRequestURI(itemBook.ImageUrl); err != nil {
		defer ErrorHandler(w)
		panic("image url harus berupa uri")
	} else if itemBook.ReleaseYear < 1980 && itemBook.ReleaseYear > 2021 {
		defer ErrorHandler(w)
		panic("release year harus antara 1980 sampai 2021")
	} else {
		fmt.Println("akan insert book")
		if err := repo.InsertBook(ctx, itemBook); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	res := map[string]string{
		"message": "berhasil tambah book",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func ErrorHandler(w http.ResponseWriter) {
	message := recover()
	res := map[string]interface{}{
		"error":    "true",
		"logError": message,
		// "message": "release year harus antara 1980 sampai 2021",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
	// return res
}

func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var itemBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&itemBook); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	//untuk mengambil matkul params
	var idBook, _ = strconv.Atoi(ps.ByName("id"))

	if itemBook.ReleaseYear >= 1980 && itemBook.ReleaseYear <= 2021 {
		if err := repo.UpdateBook(ctx, itemBook, idBook); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
	} else {
		res := map[string]string{
			"error":   "true",
			"message": "release year harus antara 1980 sampai 2021",
		}
		utils.ResponseJSON(w, res, http.StatusBadRequest)
	}

	res := map[string]string{
		"message": fmt.Sprintf("update book %d sukses", idBook),
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idBook, _ = strconv.Atoi(ps.ByName("id"))
	if err := repo.DeleteBook(ctx, idBook); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"message": fmt.Sprintf("delete book %d sukses", idBook),
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

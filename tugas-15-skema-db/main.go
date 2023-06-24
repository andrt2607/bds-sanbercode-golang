package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	config "tugas15/Config"
	"tugas15/models"
	"tugas15/repo"
	"tugas15/utils"

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

	// fmt.Println("Success")

	router := httprouter.New()
	// router.
	router.GET("/nilai", GetAllNilai)
	router.POST("/nilai", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)
	router.GET("/matkul", GetMataKuliah)
	router.POST("/matkul", PostMataKuliah)
	router.PUT("/matkul/:id/update", UpdateMataKuliah)
	router.DELETE("/matkul/:id/delete", DeleteMatkul)
	router.GET("/mahasiswa", GetAllMahasiswa)
	router.POST("/mahasiswa", PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)
	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetAllNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilais, err := repo.GetAllNilai(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, nilais, http.StatusOK)
}

func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var itemNilai models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&itemNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repo.Insert(ctx, itemNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "berhasil tambah nilai",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var nilai models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	//untuk mengambil nilai params
	var idNilai, _ = strconv.Atoi(ps.ByName("id"))

	if err := repo.UpdateNilai(ctx, nilai, idNilai); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "update success",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idNilai = ps.ByName("id")
	if err := repo.DeleteNilai(ctx, idNilai); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "berhasil delete",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func GetMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mataKuliahs, err := repo.GetAllMataKuliah(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mataKuliahs, http.StatusOK)
}

func PostMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var itemMataKuliah models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&itemMataKuliah); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repo.InsertMatKul(ctx, itemMataKuliah); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "berhasil tambah mata kuliah",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var matkul models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&matkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	//untuk mengambil matkul params
	var idMatkul, _ = strconv.Atoi(ps.ByName("id"))

	if err := repo.UpdateMatkul(ctx, matkul, idMatkul); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "update matkul sukses",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteMatkul(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMatkul = ps.ByName("id")
	if err := repo.DeleteMatkul(ctx, idMatkul); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "berhasil delete matkul",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

func GetAllMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mahasiswas, err := repo.GetAllMahasiswa(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, mahasiswas, http.StatusOK)
}

func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var itemMahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&itemMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := repo.InsertMahasiswa(ctx, itemMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "berhasil tambah mahasiswa",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	//untuk mengambil matkul params
	var idMahasiswa, _ = strconv.Atoi(ps.ByName("id"))

	if err := repo.UpdateMahasiswa(ctx, mahasiswa, idMahasiswa); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "update mahasiswa sukses",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMahasiswa = ps.ByName("id")
	if err := repo.DeleteMahasiswa(ctx, idMahasiswa); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"status": "berhasil delete mahasiswa",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
}

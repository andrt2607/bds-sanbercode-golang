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

	fmt.Println("Success")

	router := httprouter.New()
	// router.
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa", PostNilai)
	router.PUT("/mahasiswa/:id/update", UpdateNilai)
	router.DELETE("/mahasiswa/:id/delete", DeleteNilai)
	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
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

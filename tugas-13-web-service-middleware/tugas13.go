package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func AddMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newNilai NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJson := json.NewDecoder(r.Body)
			if err := decodeJson.Decode(&newNilai); err != nil {
				log.Fatal(err)
			}
			if newNilai.Nilai > 100 {
				response := []byte(`{"message": "Data nilai maksimal 100"}`)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			switch {
			case newNilai.Nilai >= 80:
				newNilai.IndeksNilai = "A"
			case newNilai.Nilai >= 70 && newNilai.Nilai < 80:
				newNilai.IndeksNilai = "B"
			case newNilai.Nilai >= 60 && newNilai.Nilai < 70:
				newNilai.IndeksNilai = "C"
			case newNilai.Nilai >= 50 && newNilai.Nilai < 60:
				newNilai.IndeksNilai = "D"
			case newNilai.Nilai < 80:
				newNilai.IndeksNilai = "E"
			}
			newNilai.ID = uint(rand.Intn(1000))
			fmt.Println("ini value nilai :", newNilai)
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			response := []byte(`{"message": "Data nilai berhasil ditambahkan"}`)
			w.Write(response)
			fmt.Println("ini arr nilai :", nilaiNilaiMahasiswa)
		} else {
			//dari form data
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mataKuliah")
			indeksNilai := r.PostFormValue("indeksNilai")
			newNilai := NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: indeksNilai,
				Nilai:       uint(nilai),
				ID:          uint(id),
			}
			if newNilai.Nilai > 100 {
				response := []byte(`{"message": "Data nilai maksimal 100"}`)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(response)
				return
			}
			// newNilai.ID = uint(rand.Intn(1000))
			switch {
			case newNilai.Nilai >= 80:
				newNilai.IndeksNilai = "A"
			case newNilai.Nilai >= 70 && newNilai.Nilai < 80:
				newNilai.IndeksNilai = "B"
			case newNilai.Nilai >= 60 && newNilai.Nilai < 70:
				newNilai.IndeksNilai = "C"
			case newNilai.Nilai >= 50 && newNilai.Nilai < 60:
				newNilai.IndeksNilai = "D"
			case newNilai.Nilai < 80:
				newNilai.IndeksNilai = "E"
			}
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
			fmt.Println("ini nilai :", nilaiNilaiMahasiswa)
			// w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response := []byte(`{"message": "Data nilai berhasil ditambahkan"}`)
			w.Write(response)
		}
	}
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	http.Handle("/mahasiswa", Auth(http.HandlerFunc(AddMahasiswa)))
	http.HandleFunc("/getMahasiswa", GetMahasiswa)
	// http.HandleFunc("/movies", getMovies)
	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

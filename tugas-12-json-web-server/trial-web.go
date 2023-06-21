package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

type Lingkaran struct {
	jariJari float64
	tinggi   float64
}

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}
func luas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody Lingkaran
	json.NewDecoder(r.Body).Decode(&requestBody)

	countLuas := math.Phi * requestBody.jariJari * requestBody.jariJari
	fmt.Println(requestBody.jariJari)
	fmt.Println(countLuas)
	response := Response{
		Data:    countLuas,
		Message: "Hasil perhitungan luas lingkaran",
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/luas", luas)
	// http.HandleFunc("/keliling", keliling)
	// http.HandleFunc("/volum", volum)

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas14/Config"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := Config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success")

	router := httprouter.New()
	router.PO
	router.GET("/mahasiswa", GetMahasiswa)
	fmt.Println("Server running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// movies, err := movie.GetAll(ctx)

	// if err != nil {
	//   fmt.Println(err)
	// }

	// utils.ResponseJSON(w, movies, http.StatusOK)
}

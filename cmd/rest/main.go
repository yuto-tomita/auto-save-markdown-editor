package main

import (
	// "myapp/pkg/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
}

func main() {
	// db := db.ConnectDB()

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
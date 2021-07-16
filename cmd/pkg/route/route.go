package route

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func serve() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world")
}
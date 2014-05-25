package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/sample", sample).Methods("GET")
	r.HandleFunc("/", sample).Methods("GET")

	http.Handle("/", r)

	log.Println("Server started. Listening...")
	http.ListenAndServe(":3000", nil)
}

func sample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there."))
}

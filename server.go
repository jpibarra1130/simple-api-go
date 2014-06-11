package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jpibarra1130/simple-api-go/controllers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", PostsHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Server started. Listening...")
	http.ListenAndServe(":3000", nil)
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := json.Marshal(controllers.GetPosts())

	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Something bad has happened."))
		return
	}

	log.Printf("Post: %v", string(out))

	w.Write([]byte(out))
}

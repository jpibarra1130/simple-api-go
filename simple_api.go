package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Post struct {
	Id    int
	Title string
	Body  string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", postsHandler).Methods("GET")
	r.HandleFunc("/", postsHandler).Methods("GET")

	http.Handle("/", r)

	log.Println("Server started. Listening...")
	http.ListenAndServe(":3000", nil)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	post := Post{Id: 1, Title: "This is a test post", Body: "This is the body."}

	out, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	log.Printf("Post: %v", string(out))

	w.Write([]byte(out))
}

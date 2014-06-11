package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jpibarra1130/simple-api-go/controllers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", PostsHandler).Methods("GET")
	r.HandleFunc("/user/register", RegisterHandler).Methods("POST")

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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	status := controllers.RegisterUser(r.FormValue("email"), r.FormValue("password"))

	w.Write([]byte("{ status : \"" + strconv.FormatBool(status) + "\" }"))
}

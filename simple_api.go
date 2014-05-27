package main

import (
	"database/sql"
	"encoding/json"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"net/http"
)

type Post struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
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
	out, err := json.Marshal(posts())

	if err != nil {
		panic(err)
	}

	log.Printf("Post: %v", string(out))

	w.Write([]byte(out))
}

func posts() []Post {
	db, err := sql.Open("mymysql", "tcp:localhost:3306*simple_api_development/root/")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	var posts []Post
	_, err = dbmap.Select(&posts, "select id, title, body from posts order by id")

	if err != nil {
		log.Printf("Error: ", err)
		return nil
	}

	return posts
}

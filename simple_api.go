package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
	"github.com/kylelemons/go-gypsy/yaml"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"net/http"
	"path/filepath"
)

type Post struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

type DbConf struct {
	Driver string
	Url    string
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

	dbConfig := DbConfig()

	if dbConfig == nil {
		return nil
	}

	log.Printf("Connecting to %v, %v", dbConfig.Driver, dbConfig.Url)

	db, err := sql.Open(dbConfig.Driver, dbConfig.Url)

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

func DbConfig() *DbConf {
	p := "db/"
	env := "development"
	cfgFile := filepath.Join(p, "dbconf.yml")

	f, err := yaml.ReadFile(cfgFile)
	if err != nil {
		return nil
	}

	drv, err := f.Get(fmt.Sprintf("%s.driver", env))
	if err != nil {
		return nil
	}

	open, err := f.Get(fmt.Sprintf("%s.open", env))
	if err != nil {
		return nil
	}

	return &DbConf{drv, open}
}

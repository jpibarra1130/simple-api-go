package controllers

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"
	"flag"
	"github.com/coopernurse/gorp"
	"github.com/jpibarra1130/simple-api-go/models"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
)

// global options. available to any subcommands. This was taken from goose library
var flagPath = flag.String("path", "db", "folder containing db info")
var flagEnv = flag.String("env", "development", "which DB environment to use")

func GetPosts() []models.Post {
	conf, err := goose.NewDBConf(*flagPath, *flagEnv)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Printf("Connecting to %v, %v", conf.Driver.Name, conf.Driver.OpenStr)

	db, err := sql.Open(conf.Driver.Name, conf.Driver.OpenStr)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	var posts []models.Post
	_, err = dbmap.Select(&posts, "select id, title, body from posts order by id")

	if err != nil {
		log.Printf("Error: ", err)
		return nil
	}

	return posts
}

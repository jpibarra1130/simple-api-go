package controllers

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/jpibarra1130/simple-api-go/models"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
)

func RegisterUser(email string, password string) bool {
	log.Println("email ", email)

	user := models.User{Email: email, HashedPassword: models.HashPassword(password)}

	conf, err := goose.NewDBConf(*flagPath, *flagEnv)

	if err != nil {
		log.Fatal(err)
		return false
	}

	log.Printf("Connecting to %v, %v", conf.Driver.Name, conf.Driver.OpenStr)

	db, err := sql.Open(conf.Driver.Name, conf.Driver.OpenStr)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	defer dbmap.Db.Close()

	err = dbmap.Insert(user)

	if err != nil {
		log.Printf("Error: ", err)
		return false
	}

	return true
}

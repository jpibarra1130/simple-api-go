package controllers

import (
	"github.com/jpibarra1130/simple-api-go/models"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
)

func RegisterUser(email string, password string) bool {
	log.Println("email ", email)

	user := models.NewUser(email, password)

	// initialize the DbMap
	dbmap := initDb()
	defer dbmap.Db.Close()

	err := dbmap.Insert(&user)

	if err != nil {
		log.Printf("Error: ", err)
		return false
	}

	return true
}

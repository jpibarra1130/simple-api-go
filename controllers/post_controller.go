package controllers

import (
	"github.com/jpibarra1130/simple-api-go/models"
	"log"
)

func GetPosts() []models.Post {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var posts []models.Post
	_, err := dbmap.Select(&posts, "select id, title, body from posts order by id")

	if err != nil {
		log.Printf("Error: ", err)
		return nil
	}

	return posts
}

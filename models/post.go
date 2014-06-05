package models

type Post struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Body  string `db:"body"`
}

package models

type User struct {
	Id             int
	Email          string
	HashedPassword []byte
	Password       string
	CreatedAt      int64
	UpdatedAt      int64
}

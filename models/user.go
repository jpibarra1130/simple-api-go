package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"log"
)

type User struct {
	Id             int    `db:"id"`
	Email          string `db:"email"`
	HashedPassword []byte `db."hashed_password"`
	Password       string
	CreatedAt      int64 `db."created_at"`
	UpdatedAt      int64 `db."updated_at"`
}

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln("Unable to hash password.", err)
		return nil
	}

	return hashedPassword
}

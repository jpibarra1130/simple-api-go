package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	Id             int64  `db:"id"`
	Email          string `db:"email"`
	HashedPassword []byte `db:"hashed_password"`
	Password       string `db:"-"`
	CreatedAt      int64  `db:"created_at"`
	UpdatedAt      int64  `db:"updated_at"`
}

func (user *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))

	if err != nil {
		return false
	}

	return true
}

func HashPassword(password string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalln("Unable to hash password.", err)
		return nil
	}

	return hashedPassword
}

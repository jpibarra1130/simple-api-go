package models

import (
	"testing"
)

func TestUserHashPassword(t *testing.T) {
	result := HashPassword("password")

	if result != nil {
		t.Log("result is correct. ", result)
	} else {
		t.Error("result is wrong")
	}
}

func TestUserValid(t *testing.T) {
	user := User{HashedPassword: HashPassword("password")}

	result := user.ValidatePassword("password")

	if result {
		t.Log("result is correct. ", result)
	} else {
		t.Error("It should be valid")
	}
}

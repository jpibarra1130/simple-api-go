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

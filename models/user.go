package models

type User struct {
	Id             int    `db:"id"`
	Email          string `db:"email"`
	HashedPassword []byte `db."hashed_password"`
	Password       string
	CreatedAt      int64 `db."created_at"`
	UpdatedAt      int64 `db."updated_at"`
}

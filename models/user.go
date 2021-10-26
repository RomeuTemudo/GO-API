package models

type User struct {
	UserID    int    `json:"id"`
	UserEmail string `json:"email"`
	Password  []byte `json:"-"`
}

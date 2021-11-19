package models

type User struct {
	UserID    int    `json:"id" gorm:"primaryKey"`
	UserEmail string `json:"email"`
	Password  string `json:"password"`
	RoleID    int    `json:"role_id"`
}

type UserOutput struct {
	UserID          int    `json:"id" gorm:"primaryKey"`
	UserEmail       string `json:"email"`
	RoleID          int    `json:"role_id"`
	RoleDescription string `json:"role_description"`
}

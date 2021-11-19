package models

type UsersRoles struct {
	RoleID          int    ` json:"id"` //key do param que vem do frontend
	RoleDescription string `json:"description"`
}

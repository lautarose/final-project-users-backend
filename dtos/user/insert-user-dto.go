package dtos

type InsertUserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
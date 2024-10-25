package models

type AuthUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

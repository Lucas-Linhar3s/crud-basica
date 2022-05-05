package models

type Usuario struct {
	Id    int32  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var Usuarios []Usuario
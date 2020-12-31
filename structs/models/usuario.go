package models

type Usuario struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	*Persona
}

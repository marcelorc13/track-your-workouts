package models

type Usuario struct {
	ID           int    `json:"id"`
	NomeCompleto string `json:"nome_completo" validate:"required,min=5,max=100"`
	Username     string `json:"username" validate:"required,min=3,max=30"`
	Email        string `json:"email" validate:"required,email"`
	Senha        string `json:"senha" validate:"required,min=6,max=30"`
}
type LoginUsuario struct {
	Email string `json:"email" validate:"required,email"`
	Senha string `json:"senha" validate:"required,min=6,max=30"`
}

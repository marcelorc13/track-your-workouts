package models

type Usuario struct {
	ID           int    `json:"id"`
	NomeCompleto string `json:"nome_completo"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Senha        string `json:"senha"`
}

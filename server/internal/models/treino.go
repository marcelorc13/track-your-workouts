package models

type (
	Treino struct {
		Nome       string
		Exercicios []Exercicio
	}

	Exercicio struct {
		Nome   string
		Series int
	}
)

package models

type Treino struct {
	ID         int
	Nome       string
	Exercicios []Exercicio
}

type Exercicio struct {
	ID     int
	Nome   string
	Series int
}

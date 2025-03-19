package domain

import "server/internal/models"

type TreinoRepository interface {
	GetTreinos() ([]models.Treino, error)
	GetTreino(id int) (models.Treino, error)
	GetExercicio(id int) (models.Exercicio, error)
}

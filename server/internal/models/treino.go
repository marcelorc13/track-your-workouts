package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Treino struct {
		ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Nome       string
		Exercicios []Exercicio
	}

	Exercicio struct {
		Nome   string
		Series int
	}
)

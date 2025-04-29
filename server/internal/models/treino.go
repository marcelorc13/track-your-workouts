package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Treino struct {
		ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Nome       string             `json:"nome"`
		CriadoPor  string             `json:"criado_por" bson:"criado_por,omitempty"`
		Exercicios []Exercicio        `json:"exercicios"`
	}

	Exercicio struct {
		Nome   string `json:"nome"`
		Series int    `json:"series"`
	}
)

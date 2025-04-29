package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Treino struct {
		ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Nome       string             `json:"nome"`
		CriadoPor  uuid.UUID          `json:"criado_por"`
		Exercicios []Exercicio        `json:"exercicios"`
	}

	Exercicio struct {
		Nome   string `json:"nome"`
		Series int    `json:"series"`
	}
)

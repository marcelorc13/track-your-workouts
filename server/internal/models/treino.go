package models

import (
	"time"

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

	//----------------------------------------------------------------

	Secao struct {
		ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		IDTreino  primitive.ObjectID `json:"id_treino" bson:"id_treino" validate:"required"`
		IDUsuario string             `json:"id_usuario" bson:"id_usuario" validate:"required"`
		Nome      string             `json:"nome" validate:"required"`
		Data      time.Time          `json:"data" validate:"required"`
		Series    []SerieExercicio   `json:"series" validate:"required"`
	}

	Serie struct {
		Repeticoes int     `json:"repeticoes"`
		Peso       float32 `json:"peso"`
	}
	SerieExercicio struct {
		Nome   string  `json:"nome"`
		Series []Serie `json:"series"`
	}
)

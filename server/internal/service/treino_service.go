package service

import (
	"fmt"
	"server/internal/models"
	"server/internal/repository"

	"github.com/go-playground/validator"
)

type TreinoService struct {
	repository repository.TreinoRepository
}

func NewTreinoService(tr repository.TreinoRepository) *TreinoService {
	return &TreinoService{tr}
}

func (ts TreinoService) CreateTreino(t models.Treino) error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	res, err := ts.repository.CreateTreino(t)

	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf(res.Message)
	}

	return nil
}

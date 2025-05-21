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

func (ts TreinoService) GetTreinosDoUsuario(usuarioId string) (*[]models.Treino, error) {
	res, err := ts.repository.GetTreinosDoUsuario(usuarioId)
	if err != nil {
		return nil, err
	}
	treinos, ok := res.Data.([]models.Treino)
	if !ok {
		return nil, fmt.Errorf("erro ao converter dados")
	}

	return &treinos, nil
}

func (ts TreinoService) GetTreinoById(id string) (*models.Treino, error) {
	res, err := ts.repository.GetTreinoById(id)
	if err != nil {
		return nil, err
	}

	treino, ok := res.Data.(models.Treino)
	if !ok {
		return nil, fmt.Errorf("erro ao converter dados")
	}

	return &treino, nil
}

func (ts TreinoService) CreateSecao(s models.Secao) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	res, err := ts.repository.CreateSecao(s)

	if err != nil {
		return fmt.Errorf(res.Message)
	}

	if !res.Success {
		return fmt.Errorf(res.Message)
	}

	return nil
}

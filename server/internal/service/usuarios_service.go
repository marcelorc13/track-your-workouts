package service

import (
	"server/internal/models"
	"server/internal/repository"
)

type UserService struct {
	repository repository.UserRepository
}

type responseError struct {
	message string
}

func (e *responseError) Error() string {
	return e.message
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{repository: userRepository}
}

func (s UserService) GetUsuarios() ([]*models.Usuario, error) {
	res, err := s.repository.GetUsuarios()

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &responseError{res.Message}
	}

	usuarios, ok := res.Data.([]*models.Usuario)

	if !ok {
		return nil, &responseError{"Erro ao converter dados"}
	}

	return usuarios, nil
}

func (s UserService) GetUsuario(id int) (*models.Usuario, error) {
	res, err := s.repository.GetUsuario(id)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &responseError{res.Message}
	}

	usuario, ok := res.Data.(models.Usuario)

	if !ok {
		return nil, &responseError{"Erro ao converter dados"}
	}

	return &usuario, nil
}

func (s UserService) DeleteUsuario(id int) error {
	res, err := s.repository.DeleteUsuario(id)

	if err != nil {
		return err
	}

	if !res.Success {
		return &responseError{res.Message}
	}

	return nil
}

package service

import (
	"server/internal/models"
	"server/internal/repository"

	"github.com/go-playground/validator"
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

func (us UserService) GetUsuarios() (*[]models.Usuario, error) {
	res, err := us.repository.GetUsuarios()

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &responseError{res.Message}
	}

	usuarios, ok := res.Data.([]models.Usuario)

	if !ok {
		return nil, &responseError{"Erro ao converter dados"}
	}

	return &usuarios, nil
}

func (us UserService) GetUsuario(id int) (*models.Usuario, error) {
	res, err := us.repository.GetUsuario(id)

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

func (us UserService) DeleteUsuario(id int) (*string, error) {
	res, err := us.repository.DeleteUsuario(id)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &responseError{res.Message}
	}

	return &res.Message, nil
}

func (us UserService) CreateUsuario(u models.Usuario) (*string, error) {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return nil, err.(validator.ValidationErrors)
	}

	res, err := us.repository.CreateUsuario(u)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, &responseError{res.Message}
	}

	return &res.Message, nil
}

package service

import (
	"fmt"
	"server/internal/models"
	"server/internal/repository"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type UserService struct {
	repository repository.UserRepository
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
		return nil, nil
	}

	usuarios, ok := res.Data.([]models.Usuario)

	if !ok {
		return nil, fmt.Errorf("erro ao converter dados")
	}

	return &usuarios, nil
}

func (us UserService) GetUsuario(id int) (*models.Usuario, error) {
	res, err := us.repository.GetUsuario(id)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, nil
	}

	usuario, ok := res.Data.(models.Usuario)

	if !ok {
		return nil, fmt.Errorf("erro ao converter dados")
	}

	return &usuario, nil
}

func (us UserService) DeleteUsuario(id int) error {
	res, err := us.repository.DeleteUsuario(id)

	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf(res.Message)
	}

	return nil
}

func (us UserService) CreateUsuario(u models.Usuario) error {
	validate := validator.New()
	u.ID = uuid.New()
	err := validate.Struct(u)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	res, err := us.repository.CreateUsuario(u)

	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf(res.Message)
	}

	return nil
}

func (us UserService) Login(u models.LoginUsuario) error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	res, err := us.repository.Login(u)

	if err != nil {
		return err
	}

	if !res.Success {
		return fmt.Errorf(res.Message)
	}

	return nil
}

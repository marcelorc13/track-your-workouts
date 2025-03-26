package repository

import (
	"database/sql"
	"fmt"
	"server/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUsuarios() (models.DBResponse, error) {

	results, err := r.DB.Query("SELECT id, nome_completo, username, email, senha FROM usuarios")

	if err != nil {
		return models.DBResponse{Success: true, Message: "Ocorreu um erro na query"}, err
	}

	defer results.Close()

	res := []models.Usuario{}

	for results.Next() {
		var user models.Usuario

		err = results.Scan(&user.ID, &user.NomeCompleto, &user.Username, &user.Email, &user.Senha)

		if err != nil {
			panic(err)
		}

		res = append(res, user)
	}

	if len(res) == 0 {
		return models.DBResponse{Success: true, Message: "O banco ainda não possui usuários"}, nil
	}

	return models.DBResponse{Success: true, Message: "Lista de todos os usuários do banco", Data: res}, nil
}
func (r *UserRepository) GetUsuario(id int) (models.DBResponse, error) {
	var user models.Usuario

	err := r.DB.QueryRow("SELECT id, nome_completo, username, email, senha FROM usuarios WHERE id = ?", id).
		Scan(&user.ID, &user.NomeCompleto, &user.Username, &user.Email, &user.Senha)

	if err == sql.ErrNoRows {
		return models.DBResponse{Message: "Usuário não encontrado"}, err
	} else if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}
	return models.DBResponse{Success: true, Message: fmt.Sprintf("Usuário de id %d encontrado", id), Data: user}, nil
}

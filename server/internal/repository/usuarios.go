package repository

import (
	"database/sql"
	"server/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUsuarios() (models.Response, error) {

	results, err := r.DB.Query("SELECT id, nome_completo, username, email, senha FROM usuarios")

	if err != nil {
		return models.Response{Sucesso: true, Status: 500, Mensagem: "Ocorreu um erro na query"}, err
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
		return models.Response{Sucesso: true, Status: 404, Mensagem: "O banco ainda não possui usuários"}, nil
	}

	return models.Response{Sucesso: true, Status: 200, Mensagem: "Lista de todos os usuários do banco", Data: res}, nil
}
func (r *UserRepository) GetUsuario(id int) (models.Response, error) {
	var user models.Usuario

	err := r.DB.QueryRow("SELECT id, nome_completo, username, email, senha FROM usuarios WHERE id = ?", id).
		Scan(&user.ID, &user.NomeCompleto, &user.Username, &user.Email, &user.Senha)

	if err == sql.ErrNoRows {
		return models.Response{Status: 404, Mensagem: "Usuário não encontrado"}, err
	} else if err != nil {
		return models.Response{Status: 500, Mensagem: err.Error()}, err
	}
	return models.Response{Sucesso: true, Status: 200, Mensagem: "Usuário de id" + string(rune(id)) + " encontrado", Data: user}, nil
}

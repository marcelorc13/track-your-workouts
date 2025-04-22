package repository

import (
	"database/sql"
	"server/internal/models"

	"golang.org/x/crypto/bcrypt"
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
		return models.DBResponse{Message: "Ocorreu um erro na query"}, err
	}

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
		return models.DBResponse{Message: "O banco ainda não possui usuários"}, nil
	}

	return models.DBResponse{Success: true, Data: res}, nil
}

func (r *UserRepository) GetUsuario(id int) (models.DBResponse, error) {
	var user models.Usuario

	err := r.DB.QueryRow("SELECT id, nome_completo, username, email, senha FROM usuarios WHERE id = ?", id).
		Scan(&user.ID, &user.NomeCompleto, &user.Username, &user.Email, &user.Senha)

	if err == sql.ErrNoRows {
		return models.DBResponse{Message: "usuário não encontrado"}, nil
	} else if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}
	return models.DBResponse{Success: true, Data: user}, nil
}

func (r *UserRepository) DeleteUsuario(id int) (models.DBResponse, error) {
	res, err := r.DB.Exec("DELETE FROM usuarios WHERE id = ?", id)

	if err != nil {
		return models.DBResponse{Message: "ocorreu um erro na query"}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return models.DBResponse{Message: "ocorreu um erro na query"}, err
	}

	if rows != 1 {
		return models.DBResponse{Message: "usuário não encontrado"}, nil
	}

	return models.DBResponse{Success: true}, nil
}

func (r *UserRepository) CreateUsuario(u models.Usuario) (models.DBResponse, error) {
	senhaBytes, err := bcrypt.GenerateFromPassword([]byte(u.Senha), 14)

	if err != nil {
		return models.DBResponse{Message: "erro ao hashear senha"}, nil
	}

	res, err := r.DB.Exec(`
		INSERT INTO usuarios(nome_completo, username, email, senha)
		VALUES(?, ?, ?, ?);
	`, u.NomeCompleto, u.Username, u.Email, string(senhaBytes))

	if err != nil {
		return models.DBResponse{Message: "erro ao criar usuario"}, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return models.DBResponse{Message: "ocorreu um erro na query"}, err
	}

	if rows != 1 {
		return models.DBResponse{Message: "erro ao criar usuario"}, err
	}

	return models.DBResponse{Success: true}, nil
}

func (r *UserRepository) Login(u models.LoginUsuario) (models.DBResponse, error) {
	var usuario models.LoginUsuario
	err := r.DB.QueryRow("SELECT email, senha FROM usuarios WHERE email = ?", u.Email).
		Scan(&usuario.Email, &usuario.Senha)

	if err == sql.ErrNoRows {
		return models.DBResponse{Message: "usuário não encontrado"}, nil
	} else if err != nil {
		return models.DBResponse{Message: err.Error()}, err
	}

	errSenha := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(u.Senha))

	if errSenha != nil {
		return models.DBResponse{Message: "senha incorreta"}, nil
	}
	return models.DBResponse{Success: true}, nil
}

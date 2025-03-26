package main

import (
	"fmt"
	"server/internal/repository"
	"server/pkg/database"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	treinoRepo := repository.NewUserRepository(db)

	usuarios, err := treinoRepo.GetUsuarios()
	if err != nil {
		fmt.Println(usuarios.Message)
	}
	fmt.Println(usuarios.Data)

	// usuario, err := treinoRepo.GetUsuario(2)
	// if err != nil {
	// 	fmt.Println(usuario.Message)
	// }
	// fmt.Println(usuario.Message, usuario.Data)

	// res, err := treinoRepo.DeleteUsuario(5)
	// if err != nil {
	// 	fmt.Println(res.Message)
	// }
	// fmt.Println(res.Message)

	// usuarioTeste := models.Usuario{NomeCompleto: "nome teste", Username: "teste1", Email: "emailteste14432@gmail.com", Senha: "123144"}
	// create, err := treinoRepo.CreateUsuario(usuarioTeste)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(create.Message, create.Data)
}

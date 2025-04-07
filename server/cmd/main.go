package main

import (
	"fmt"
	"server/internal/repository"
	"server/pkg/database"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	// userServ := service.NewUserService(*userRepo)

	// usuarios, err := userRepo.GetUsuarios()
	// if err != nil {
	// 	fmt.Println(usuarios.Message)
	// }
	// fmt.Println(usuarios.Data)

	usuario, err := userRepo.GetUsuario(2)
	if err != nil {
		fmt.Println(usuario.Message)
	}
	fmt.Println(usuario.Message, usuario.Data)

	// res, err := userRepo.DeleteUsuario(2)
	// if err != nil {
	// 	fmt.Println(res.Message)
	// }
	// fmt.Println(res.Message)

	// usuarioTeste := models.Usuario{NomeCompleto: "nome teste", Username: "tes1ads", Email: "emailtdse4432@gmail.com", Senha: "123144"}
	// create, err := userRepo.CreateUsuario(usuarioTeste)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(create.Message, create.Data)
}

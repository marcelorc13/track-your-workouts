package main

import (
	"fmt"
	"server/internal/repository"
	"server/internal/service"
	"server/pkg/database"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userServ := service.NewUserService(*userRepo)

	usuarios, err := userServ.GetUsuarios()
	if err != nil {
		fmt.Println(err)
	}
	for _, u := range *usuarios {
		fmt.Println(u)
	}

	// usuario, err := userServ.GetUsuario(6)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*usuario)

	// res, err := userRepo.DeleteUsuario(2)
	// if err != nil {
	// 	fmt.Println(res.Message)
	// }
	// fmt.Println(res.Message)

	// usuarioTeste := models.Usuario{NomeCompleto: "nome teste", Username: "kdsdmalkd", Email: "emailtdsesdds4432¨@gmail.com", Senha: "123144"}
	// create, err := userServ.CreateUsuario(usuarioTeste)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*create)
}

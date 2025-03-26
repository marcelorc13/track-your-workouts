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
		fmt.Println(err)
	}
	fmt.Println(usuarios)

	usuario, err := treinoRepo.GetUsuario(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usuario)
}

package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	godotenv.Load()

	usuario := os.Getenv("MYSQL_USER")
	senha := os.Getenv("MYSQL_PASSWORD")
	endereco := os.Getenv("MYSQL_ADDRESS")
	database := os.Getenv("MYSQL_DB")

	// dsn := "user:pass@tcp(127.0.0.1:3306)/teste"
	dsn := usuario + ":" + senha + "@tcp(" + endereco + ")/" + database

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	return db

}

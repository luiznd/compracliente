package config

import (
	"database/sql"
	"fmt"

	"../fileinput"
	_ "github.com/lib/pq"
)

//Disponibiliza conexão
func GetPostgresDB() (*sql.DB, error) {

	user := "-"
	password := "-"
	host := "-"
	port := "-"
	dbname := "-"

	var conteudo []string
	conteudo, err := fileinput.GetText("./files/config.ini")
	if err != nil {
		fmt.Println(err)
	}

	// Atribui os parametros de conexão postgres contidos no arquivo ../files/config.ini
	for indice := range conteudo {
		if indice == 0 {
			user = conteudo[indice]
		} else if indice == 1 {
			password = conteudo[indice]
		} else if indice == 2 {
			host = conteudo[indice]
		} else if indice == 3 {
			port = conteudo[indice]
		} else if indice == 4 {
			dbname = conteudo[indice]
		}
	}

	desc := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

//Cria conexão
func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}

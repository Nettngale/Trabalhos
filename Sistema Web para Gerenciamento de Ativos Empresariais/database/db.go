package database

import (
	"database/sql"
	"fmt"
)

func ConectaBanco() *sql.DB {

	//Configuração da Conexão - usuario:senha(ip:porta)/base
	conexao := "root:nett@tcp(127.0.0.1:3306)/lista_ativos"
	//abrir conexão - informe o driver e a string de conexão
	db, err := sql.Open("mysql", conexao)

	if err != nil {
		fmt.Print("Erro ao conectar a lista: ", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("Erro ao buscar o servidor", err)
	} else {
		fmt.Print("Conexão realizada com sucesso!!")
	}
	return db
}

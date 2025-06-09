package service

import (
	"fmt"
	"projeto/database"
)

type Ativos struct {
	Id   int
	Nome string
}

const insertComando string = "insert into ativos(nome_ativos) values (?)"
const selectComando string = "select id_ativos, nome_ativos from ativos"

func InsereAtivo(nome string) {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Exec(insertComando, nome)
	if erro != nil {
		fmt.Print("Erro ao inserir doador", erro)
	} else {
		fmt.Print(resultado)
	}
}

func SelecionaTodosAtivos() []Ativos {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Query(selectComando)
	if erro != nil {
		fmt.Print("Erro ao buscar do banco", erro)
		return nil
	}
	var ativos []Ativos
	for resultado.Next() {
		var ativo Ativos
		err := resultado.Scan(&ativo.Id, &ativo.Nome)
		if err != nil {
			fmt.Println("Erro ao ler a linha", err)
			continue
		}
		ativos = append(ativos, ativo)
	}
	fmt.Print(ativos)
	return ativos
}

package service

import (
	"fmt"
	"projeto/database"

)

type Cliente struct {
	Id       int
	Nome     string
	Telefone string
	DataN    string
	CEP		 string
	Renda    string
}

const insertComando string = "insert into cliente(nome_cliente, telefone_cliente, dataN_cliente, cep_cliente, renda_cliente) values (?, ?, ?, ?)"
const selectComando string = "select id_cliente, nome_cliente, telefone_cliente, cep_cliente, renda_cliente, dataN from cliente"

func InsereCliente(nome string, telefone string, datan string, cep string, renda string) {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Exec(insertComando, nome, telefone, datan, cep, renda)
	if erro != nil {
		fmt.Print("Erro ao inserir doador", erro)
	} else {
		fmt.Print(resultado)
	}
}

func SelecionaTodosClientes() []Cliente {
	conexao := database.ConectaBanco()
	resultado, erro := conexao.Query(selectComando)
	if erro != nil {
		fmt.Print("Erro ao buscar do banco", erro)
		return nil
	}
	var clientes []Cliente
	for resultado.Next() {
		var cliente Cliente
		err := resultado.Scan(&cliente.Id, &cliente.Nome, &cliente.Telefone, &cliente.CEP, &cliente.DataN, &cliente.Renda)
		if err != nil {
			fmt.Println("Erro ao ler a linha", err)
			continue
		}
		clientes = append(clientes, cliente)
	}
	fmt.Print(clientes)
	return clientes
}
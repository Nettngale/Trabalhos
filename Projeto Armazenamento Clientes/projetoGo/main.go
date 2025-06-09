package main

import (
	"fmt"
	"html/template"
	"net/http"
	"projeto/database"
)

func abreIndex(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/index.html")
	if erro != nil {
		fmt.Println("O erro foi", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/cadastro.html")
	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}
	pagina.Execute(resposta, nil)
}

func salvaDoador(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/cadastro.html")
	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}

	if requisicao.Method == http.MethodPost {
		nomeDoador := requisicao.FormValue("nome")
		telefoneDoador := requisicao.FormValue("tel")
		emailDoador := requisicao.FormValue("email")
		tipoDoador := requisicao.FormValue("tipo")

		fmt.Println("O nome do Doador é: ", nomeDoador)
		fmt.Println("O Telefone do Doador é: ", telefoneDoador)
		fmt.Println("O Email do Doador é: ", emailDoador)
		fmt.Println("O Tipo do Doador é: ", tipoDoador)
		database.ConectaBanco()
	}

	pagina.Execute(resposta, nil)
}

func main() {
	//Criando um EndPoint
	http.HandleFunc("/inicial", abreIndex)
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/salvadoador", salvaDoador)

	fmt.Print("Iniciando Servidor!!!!")
	//Subindo o servidor na porta 8080
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Print("Servidor com Problemas")
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"projeto/database"
)

/*Funçao para abrir a pagina principal (Home)*/
func abreHome(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/home.html")

	/*Condiçao para mostrar se ocorreu um erro e o tipo do erro*/
	if erro != nil {
		fmt.Println("O erro foi", erro)
		return
	}
	/*Caso nao houver erro a pagina ira executar normalmente*/
	pagina.Execute(resposta, nil)
}

func abreLista(resposta http.ResponseWriter, requisicao *http.Request) {
	pagina, erro := template.ParseFiles("template/lista.html")

	if erro != nil {
		fmt.Println("O erro foi", erro)
		return
	}
	/*Caso nao houver erro a pagina ira executar normalmente*/
	pagina.Execute(resposta, nil)

}

func salvaAtivo(resposta http.ResponseWriter, requisicao *http.Request) {

	/*Variavel pagina e erro. Carrega o arquivo cadastro.html e se ocorrer
	algum erro, ele é armazenado na variavel erro.*/
	pagina, erro := template.ParseFiles("template/lista.html")

	/*Se houver um erro, erro!= nil "erro diferente de um valor nulo",
	o que foi armazenado na variavel erro sera indicado aqui.*/
	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}

	/*Verificaçao para armazenar os dados de um cliente.*/
	if requisicao.Method == http.MethodPost {

		/*Variaveis para pegar dados do cliente como: cep, renda, nome.
		Baseado nos termos do arquivo html, no caso, cadastro.html.
		Esses termos ("nome", "tel", etc) devem corresponder ao
		formulario feito no arquivo cadastro.html*/
		nomeativo := requisicao.FormValue("nome")

		/*Responsavel por imprimir os dados recebidos do usuario*/
		fmt.Println("O nome do Ativo é: ", nomeativo)

		/*Funçao para conectar ao banco de dados e armazenar as
		informaçoes recebidas.*/
		database.ConectaBanco()
	}
	/*Envia a resposta ao usuario. Representa a resposta
	HTTP (http.ResponseWriter) enviada ao navegador */
	pagina.Execute(resposta, nil)
}

func main() {
	//Criando um EndPoint
	http.HandleFunc("/home", abreHome)
	http.HandleFunc("/lista", abreLista)

	fmt.Print("Iniciando Servidor!!!!")

	erro := http.ListenAndServe(":8080", nil)
	/*Condiçao para mostrar se ocorreu um erro e o tipo do erro*/
	if erro != nil {
		fmt.Print("Servidor com Problemas")
	}
}

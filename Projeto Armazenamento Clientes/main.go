package main

import (
	"fmt"
	"html/template"
	"net/http"
	"projeto/database"
)

/*
Funçao para "abrir" o arquivo index.html. Envia a resposta
ao navegador e faz a requisição feita pelo usuário
*/
func abreIndex(resposta http.ResponseWriter, requisicao *http.Request) {

	/*Variavel pagina e erro. Carrega o arquivo index.html e se ocorrer
	algum erro, ele é armazenado na variavel erro.*/
	pagina, erro := template.ParseFiles("templates/index.html")

	/*Se houver um erro, erro!= nil "erro diferente de um valor nulo",
	o que foi armazenado na variavel erro sera indicado aqui.*/
	if erro != nil {
		fmt.Println("O erro foi", erro)
		return
	}

	/*Caso nao houver erro, a pagina sera executada normalmente
	e nil irá garantir que nada alem fique aparente na pagina */
	pagina.Execute(resposta, nil)
}

/*
Funçao para "abrir" o arquivo cadastro.html. Envia a resposta
ao navegador e faz a requisição feita pelo usuário
*/
func abreCadastro(resposta http.ResponseWriter, requisicao *http.Request) {

	/*Variavel pagina e erro. Carrega o arquivo cadastro.html e se ocorrer
	algum erro, ele é armazenado na variavel erro.*/
	pagina, erro := template.ParseFiles("templates/cadastro.html")

	/*Se houver um erro, erro!= nil "erro diferente de um valor nulo",
	o que foi armazenado na variavel erro sera indicado aqui.*/
	if erro != nil {
		fmt.Println("Erro em cadastro", erro)
		return
	}

	/*Caso nao houver erro, a pagina sera executada normalmente
	e nil irá garantir que nada alem fique aparente na pagina */
	pagina.Execute(resposta, nil)
}

/*Funçao para salvar os dados do cliente*/
func salvaCliente(resposta http.ResponseWriter, requisicao *http.Request) {

	/*Variavel pagina e erro. Carrega o arquivo cadastro.html e se ocorrer
	algum erro, ele é armazenado na variavel erro.*/
	pagina, erro := template.ParseFiles("template/cadastro.html")

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
		nomecliente := requisicao.FormValue("nome")
		telefonecliente := requisicao.FormValue("tel")
		rendacliente := requisicao.FormValue("renda")
		nascimentocliente := requisicao.FormValue("datanascimento")
		cepcliente := requisicao.FormValue("cep")

		/*Responsavel por imprimir os dados recebidos do usuario*/
		fmt.Println("O nome do Cliente é: ", nomecliente)
		fmt.Println("O Telefone do Cliente é: ", telefonecliente)
		fmt.Println("O Email do Cliente é: ", rendacliente)
		fmt.Println("O Tipo do Cliente é: ", nascimentocliente)
		fmt.Println("O Tipo do Cliente é: ", cepcliente)

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
	http.HandleFunc("/inicial", abreIndex)
	http.HandleFunc("/cadastro", abreCadastro)
	http.HandleFunc("/salvacliente", salvaCliente)
	database.ConectaBanco()
	fmt.Print("Iniciando Servidor!!!!")
	//Subindo o servidor na porta 8080
	erro := http.ListenAndServe(":8080", nil)
	if erro != nil {
		fmt.Print("Servidor com Problemas")
	}
}

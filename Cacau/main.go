package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Handle("/compra", IsAuthenticated(http.HandlerFunc(comprarHandler)))
	http.HandleFunc("/api/verificar-login", VerificarLoginHandler)
	http.HandleFunc("/cadastrar", CadastrarHandler)
	http.HandleFunc("/perfil", PerfilHandler)
	http.HandleFunc("/comprar", comprarHandler)

	// Adicione outras rotas aqui
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	db, err := sql.Open("mysql", "root:aluno@tcp(localhost:3306)/cacau_show")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados!")

}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := pegarTokenDoCookieOuHeader(r)
		if !tokenEhValido(token) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func VerificarLoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("sessao")
	if err != nil || !sessaoEhValida(cookie.Value) {
		http.Error(w, "Não autorizado", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func comprarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := r.FormValue("nome")
	endereco := r.FormValue("endereco")
	cartao := r.FormValue("cartao_credito")
	cvv := r.FormValue("cvv")
	produto := r.FormValue("produto") // novo campo produto
	usuario := r.FormValue("usuario")
	senha := r.FormValue("senha")

	db, err := sql.Open("mysql", "root:aluno@tcp(127.0.0.1:3306)/cacau_show")
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Agora insira o produto também
	_, err = db.Exec(
		"INSERT INTO clientes (nome, endereco, cartao_credito, cvv, produto, usuario, senha) VALUES (?, ?, ?, ?, ?, ?, ?)",
		nome, endereco, cartao, cvv, produto, usuario, senha,
	)
	if err != nil {
		http.Error(w, "Erro ao salvar no banco", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Compra do %s registrada com sucesso!", produto)

}

func sessaoEhValida(sessao string) bool {
	db, err := sql.Open("mysql", "root:aluno@tcp(127.0.0.1:3306)/cacau_show")
	if err != nil {
		fmt.Println("Erro ao conectar no banco:", err)
		return false
	}
	defer db.Close()

	var usuario string
	err = db.QueryRow("SELECT usuario FROM clientes WHERE usuario = ?", sessao).Scan(&usuario)
	if err != nil {
		return false
	}

	return true
}

func pegarTokenDoCookieOuHeader(r *http.Request) string {
	cookie, err := r.Cookie("sessao")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func tokenEhValido(token string) bool {
	return token != "" && sessaoEhValida(token)
}

func CadastrarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	usuario := r.FormValue("usuario")
	senha := r.FormValue("senha")
	repsenha := r.FormValue("repsenha")

	if usuario == "" || senha == "" || repsenha == "" {
		http.Error(w, "Todos os campos são obrigatórios", http.StatusBadRequest)
		return
	}

	if senha != repsenha {
		http.Error(w, "As senhas não coincidem", http.StatusBadRequest)
		return
	}

	// Conexão com o banco
	db, err := sql.Open("mysql", "root:aluno@tcp(127.0.0.1:3306)/cacau_show")
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Verificar se o usuário já existe
	var existe int
	err = db.QueryRow("SELECT COUNT(*) FROM clientes WHERE usuario = ?", usuario).Scan(&existe)
	if err != nil {
		http.Error(w, "Erro ao verificar usuário", http.StatusInternalServerError)
		return
	}

	if existe > 0 {
		http.Error(w, "Usuário já cadastrado", http.StatusConflict)
		return
	}

	// Inserir novo usuário
	_, err = db.Exec("INSERT INTO clientes (usuario, senha) VALUES (?, ?)", usuario, senha)
	if err != nil {
		http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "sessao",
		Value:    usuario, // ou um token, se for usar um
		Path:     "/",
		HttpOnly: true,
	})

	// Redirecionar para login
	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}

type Usuario struct {
	Nome     string
	Endereco string
}

func PerfilHandler(w http.ResponseWriter, r *http.Request) {
	// Exemplo: pegar o nome e endereço do usuário da sessão/banco.
	// Aqui você deve substituir pela sua lógica real:
	usuario := Usuario{
		Nome:     "João da Silva",
		Endereco: "Rua das Flores, 123 - São Paulo",
	}

	tmpl, err := template.ParseFiles("templates/perfil.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, usuario)
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		return
	}
}

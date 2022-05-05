package controllers

import (
	"crud/database"
	"crud/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> OLÁ MUNDO </h1>")
}

//INSERIR USUARIOS NO BANCO DE DADOS;
func RotaUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	requisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha no corpo da requisicao!"))
		return
	}

	var usuarios models.Usuario

	if err = json.Unmarshal(requisicao, &usuarios); err != nil {
		w.Write([]byte("Falha ao converter usuario para atruct!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao converter e conectar ao banco de dados!"))
		return
	}

	statement, err := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	insercao, err := statement.Exec(usuarios.Nome, usuarios.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement!"))
	}

	idinsercao, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido! "))
		return
	}

	w.WriteHeader(http.StatusCreated) // MUDA O RETORNO O CODIGO DA API NO POSTMAN EX 204, 404...
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idinsercao)))

}

// BUSCAR USUARIOS NO BANCO DE DADOS;
func BuscaUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
	}
	defer db.Close()

	// SELECT * FROM USUARIOS;

	linhas, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar usuarios!"))
		return
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if err := linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuario!"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter usuarios em JSON!"))
		return
	}
}

// BUSCAR USUARIO ESPECIFICO NO BANCO DE DADOS;
func BuscaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parametros := mux.Vars(r)

	ID, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro1"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados!"))
		return
	}
	defer db.Close()

	linha, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuario!"))
		return
	}

	var usuarios models.Usuario

	if linha.Next() {

		if err := linha.Scan(&usuarios.Id, &usuarios.Nome, &usuarios.Email); err != nil {
			w.Write([]byte("Erro!"))
		}
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("ERRO AO CONVERTER EM JSON"))
		return
	}
}

// ATUALIZAR O USUARIO NO BANCO DE DADOS;
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, err := strconv.ParseUint(parametro["id"], 10, 32)
	if err != nil {
		w.Write([]byte("ERRO AO LER O PARAMETRO ID!!!"))
		return
	}

	requisicao, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("ERRO NA REQUISICAO!"))
		return
	}

	var usuario models.Usuario
	if err := json.Unmarshal(requisicao, &usuario); err != nil {
		w.Write([]byte("erro no corpo da requisicao!!!"))
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("ERRO AO CONECTAR AO BANCO DE DADOS!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email= ? where id = ?")
	if err != nil {
		w.Write([]byte("erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(usuario.Nome, usuario.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar usuario!"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DELETAR USUARIOS;
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametro := mux.Vars(r)

	ID, err := strconv.ParseUint(parametro["id"], 10, 32)
	if err != nil {
		w.Write([]byte("ERRO AO CONVERTE PARAMETRO PARA INTEIRO!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("ERRO AO CONECTAR AO BANCO DE DADOS!"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("ERRO AO CRIAR STATEMENT!!!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("ERRO AO DELETAR USUARIO!!!"))
		return
	}

	w.WriteHeader(http.StatusOK)

}

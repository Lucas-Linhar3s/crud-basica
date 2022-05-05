package servidor

import (
	"fmt"
	"log"
	"net/http"
	"crud/controllers"

	"github.com/gorilla/mux"
)

func ConexaoServidor() {
	router := mux.NewRouter()

	// CRIAÇÃO DAS ROTAS

	// ROTA HOME
	router.HandleFunc("/", controllers.Home)

	// ROTA PARA CRIAR USUARIOS!
	router.HandleFunc("/usuarios", controllers.RotaUsuarios).Methods(http.MethodPost)

	// ROTA PARA BUSCAR TODOS OS USUARIOS!
	router.HandleFunc("/usuarios", controllers.BuscaUsuarios).Methods(http.MethodGet)

	// ROTA PARA BUSCAR 1 USUARIO EM ESPECIFICO!
	router.HandleFunc("/usuarios/{id}", controllers.BuscaUsuario).Methods(http.MethodGet)

	// ROTA PARA ATUALIZAR USUARIOS!
	router.HandleFunc("/usuarios/{id}", controllers.AtualizarUsuario).Methods(http.MethodPut)

	// ROTA PARA DELETAR USUARIOS!
	router.HandleFunc("/usuarios/{id}", controllers.DeletarUsuario).Methods(http.MethodDelete)

	// SERVIDOR
	fmt.Println("SERVIDOR RODANDO NA URL localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

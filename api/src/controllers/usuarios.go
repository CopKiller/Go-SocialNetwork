package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	erro = usuario.CheckDataRules()
	if erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioID, erro := repositorio.Criar(usuario)
	if erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))
}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, err := banco.Conectar()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarios, err := repositorio.BuscarTodos()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Encodar a resposta como JSON e enviar ao cliente
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(usuarios)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Logar a resposta para fins de depuração
	log.Printf("Buscar usuarios: %v", usuarios)
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando Usuário!"))
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuario!"))
}
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuario!"))
}

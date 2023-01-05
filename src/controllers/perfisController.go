package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"
	"strings"
)

func CriarPerfil(w http.ResponseWriter, r *http.Request) {
}

func BuscarPerfis(w http.ResponseWriter, r *http.Request) {

	descricao := strings.ToLower(r.URL.Query().Get("perfil"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	perfis, erro := repositorio.Buscar(descricao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, perfis)

}

func BuscarPerfil(w http.ResponseWriter, r *http.Request) {

}

func AtualizarPerfil(w http.ResponseWriter, r *http.Request) {
}

func DeletarPerfil(w http.ResponseWriter, r *http.Request) {

}

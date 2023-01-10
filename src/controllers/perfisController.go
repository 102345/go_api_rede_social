package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CriarPerfil(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var perfil modelos.Perfil
	if erro = json.Unmarshal(corpoRequest, &perfil); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = perfil.Validar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	perfidId, erro := repositorio.Criar(perfil)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	perfil.PerfilId = perfidId

	respostas.JSON(w, http.StatusCreated, perfil)
}

func BuscarPerfis(w http.ResponseWriter, r *http.Request) {

	descricao := strings.ToLower(r.URL.Query().Get("descricao"))

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

	if perfis == nil {
		respostas.Erro(w, http.StatusNotFound, errors.New("Não existe perfil informado com este parâmetro de busca!"))
		return
	}

	respostas.JSON(w, http.StatusOK, perfis)

}

func BuscarPerfil(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	perfilID, erro := strconv.ParseUint(parametros["perfilId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	perfil, erro := repositorio.BuscarPorID(perfilID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if perfil.PerfilId == 0 {
		respostas.Erro(w, http.StatusNotFound, errors.New("Não existe perfil informado!"))
		return
	}

	respostas.JSON(w, http.StatusOK, perfil)

}

func AtualizarPerfil(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	perfilID, erro := strconv.ParseUint(parametros["perfilId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var perfil modelos.Perfil
	if erro = json.Unmarshal(corpoRequisicao, &perfil); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = perfil.Validar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	if erro = repositorio.Atualizar(perfilID, perfil); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func DeletarPerfil(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	perfilID, erro := strconv.ParseUint(parametros["perfilId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	if erro = repositorio.Deletar(perfilID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func VincularUsuarioComPerfil(w http.ResponseWriter, r *http.Request) {

	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var perfilUsuario modelos.PerfilUsuario
	if erro = json.Unmarshal(corpoRequest, &perfilUsuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = perfilUsuario.Validar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	erro = repositorio.VincularPerfilDeUsuario(perfilUsuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func DeletarPerfilDeUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	perfil, erro := repositorio.BuscarPerfilDoUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if perfil.PerfilId == 0 {
		respostas.Erro(w, http.StatusNotFound, errors.New("Não existe perfil do usuário informado!"))
		return
	}

	if erro = repositorio.DeletarPerfilDeUsuario(usuarioID, perfil.PerfilId); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)

}

func BuscarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePerfis(db)
	perfil, erro := repositorio.BuscarPerfilDoUsuario(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if perfil.PerfilId == 0 {
		respostas.Erro(w, http.StatusNotFound, errors.New("Não existe perfil do usuário informado!"))
		return
	}

	respostas.JSON(w, http.StatusOK, perfil)
}

package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPerfis = []Rota{
	{
		Uri:                "/perfis",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPerfil,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfis",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPerfis,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPerfil,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPerfil,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPerfil,
		RequerAutenticacao: true,
	},
}

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
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPerfil,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPerfil,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfis/{perfilId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPerfil,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/perfis/{usuarioId}/deletar-perfil",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DeletarPerfilDeUsuario,
		RequerAutenticacao: true,
	},
}

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
		Uri:                "/perfis/vincular-usuario-perfil",
		Metodo:             http.MethodPost,
		Funcao:             controllers.VincularUsuarioComPerfil,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfis/{usuarioId}/deletar-usuario-perfil",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPerfilDeUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/perfis/{usuarioId}/buscar-usuario-perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
}

package modelos

//PerfilUsuario representa o formato do vinculo entre perfil e usuario
type PerfilUsuario struct {
	UsuarioId uint64 `json:"usuarioId"`
	PerfilId  uint64 `json:"perfilId"`
}

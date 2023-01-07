package modelos

//Perfil representa o formato do perfil do usuario
type Perfil struct {
	PerfilId  uint64 `json:"perfidId"`
	Descricao string `json:"descricao"`
}

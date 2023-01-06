package modelos

import "errors"

//Perfil representa o formato do perfil do usuario
type Perfil struct {
	PerfilId  uint64 `json:"perfidId"`
	Descricao string `json:"descricao"`
}

func (perfil *Perfil) Validar() error {

	if perfil.Descricao == "" {
		return errors.New("A descrição é obrigatória e não pode estar em branco")
	}

	return nil
}

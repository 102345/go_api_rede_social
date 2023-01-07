package modelos

import "errors"

func (perfil *Perfil) Validar() error {

	if perfil.Descricao == "" {
		return errors.New("A descrição é obrigatória e não pode estar em branco")
	}

	return nil
}

package modelos

import (
	"errors"
)

func (perfilUsuario *PerfilUsuario) Validar() error {

	if perfilUsuario.UsuarioId == 0 {
		return errors.New("O usuarioId é obrigatório e não pode estar em branco")
	}

	if perfilUsuario.PerfilId == 0 {
		return errors.New("O perfilId é obrigatório e não pode estar em branco")
	}

	return nil
}

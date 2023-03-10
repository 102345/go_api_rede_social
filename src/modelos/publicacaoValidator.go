package modelos

import (
	"errors"
	"strings"
)

// Preparar chamará os métodos validar e formatar
func (publicacao *Publicacao) Preparar() error {

	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()

	return nil
}

func (publicacao *Publicacao) validar() error {

	if publicacao.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar eme branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar eme branco")
	}

	return nil

}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}

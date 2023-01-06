package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Perfis representa um repositorio de Perfis
type Perfis struct {
	db *sql.DB
}

// NovoRepositorioDePerfis cria um repositorio de perfis
func NovoRepositorioDePerfis(db *sql.DB) *Perfis {
	return &Perfis{db}
}

// Criar insere um perfil no banco de dados
func (repositorio Perfis) Criar(perfil modelos.Perfil) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into perfis(descricao)values(?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(perfil.Descricao)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Buscar traz todos os perfis que atendem um filtro de descricao
func (repositorio Perfis) Buscar(descricao string) ([]modelos.Perfil, error) {

	descricao = fmt.Sprintf("%%%s%%", descricao)

	linhas, erro := repositorio.db.Query(
		"select perfil_id, descricao from perfis where descricao LIKE ?",
		descricao,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var perfis []modelos.Perfil

	for linhas.Next() {
		var perfil modelos.Perfil

		if erro = linhas.Scan(
			&perfil.PerfilId,
			&perfil.Descricao,
		); erro != nil {
			return nil, erro
		}

		perfis = append(perfis, perfil)
	}

	return perfis, nil
}

// BuscarPorID traz um perfil pelos seu ID
func (repositorio Perfis) BuscarPorID(ID uint64) (modelos.Perfil, error) {

	linhas, erro := repositorio.db.Query(
		"select perfil_id, descricao from perfis where perfil_id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Perfil{}, erro
	}

	defer linhas.Close()

	var perfil modelos.Perfil

	if linhas.Next() {

		if erro = linhas.Scan(
			&perfil.PerfilId,
			&perfil.Descricao,
		); erro != nil {
			return modelos.Perfil{}, erro
		}
	}

	return perfil, nil
}

// Atualizar altera as informações do perfil pelo seu ID
func (repositorio Perfis) Atualizar(ID uint64, perfil modelos.Perfil) error {

	statement, erro := repositorio.db.Prepare("update perfis set descricao = ? where perfil_id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(perfil.Descricao, ID); erro != nil {
		return erro
	}

	return nil

}

// Deletar exclui as informações do perfil
func (repositorio Perfis) Deletar(ID uint64) error {

	statement, erro := repositorio.db.Prepare("delete from perfis where perfil_id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil

}

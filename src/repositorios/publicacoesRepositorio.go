package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Publicacoes representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

//NovoRepositorioDePublicacoes cria um repositorio de publicacoes
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

//Criar insere a publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes(titulo, conteudo, autor_id)values(?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

//Buscar traz todas as publicacoes que atendem um usuario especifico e o que ele segue
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`select distinct p.*, u.nick from publicacoes p
	  				inner join usuarios u on u.id = p.autor_id 
					inner join seguidores s on p.autor_id = s.usuario_id
					where u.id = ? or s.seguidor_id = ? order by 1 desc`,
		usuarioID, usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

//BuscarPorID traz uma publicação pelo seu ID
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`select p.*, u.nick from
		      publicacoes p inner join usuarios u on p.autor_id = u.id
			  where p.id = ?`,
		publicacaoID,
	)

	if erro != nil {
		return modelos.Publicacao{}, erro
	}

	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

//Atualizar altera a publicacao do usuário pelos seu ID
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {

	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ? , conteudo = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil

}

// Deletar exclui a publicação do usuário
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil

}

//BuscarPorUsuario traz todas as publicacoes que pertencem a um usuario especifico
func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {

	linhas, erro := repositorio.db.Query(
		`select distinct p.*, u.nick from publicacoes p
	  				inner join usuarios u on u.id = p.autor_id 
					where p.autor_id = ? order by 1 desc`,
		usuarioID,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

//Curtir insere uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare(
		"update publicacoes set curtidas = curtidas + 1 where id = ?",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

//Descurtir subtrai uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {

	statement, erro := repositorio.db.Prepare(
		`update publicacoes 
		set curtidas = 
		case 
			when curtidas > 0 then curtidas - 1
			else 0 
		end
		where id = ?`,
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

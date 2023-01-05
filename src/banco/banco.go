package banco

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão DB
)

//Conectar abre a conexao do banco de dados
func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		fmt.Printf("erro de conexao : %s", config.StringConexaoBanco)
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		fmt.Printf("erro de conexao no ping : %s", config.StringConexaoBanco)
		db.Close()
		return nil, erro
	}

	return db, nil

}

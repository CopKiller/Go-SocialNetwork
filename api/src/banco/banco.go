package banco

import (
	"api/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}

package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios representa um repositório de usuarios.
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioUsuarios cria um repositório de usuarios.
func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?,?,?,?)")

	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) BuscarTodos() ([]modelos.Usuario, error) {
	query, erro := repositorio.db.Query("SELECT * FROM usuarios")
	if erro != nil {
		return nil, erro
	}
	defer query.Close()

	// Cria um slice de usuários inicialmente vazio
	usuarios := make([]modelos.Usuario, 0)

	// Itera sobre as linhas retornadas pela query
	for query.Next() {
		var usuario modelos.Usuario

		// Escaneia cada coluna do resultado para os campos correspondentes do struct Usuario
		erro = query.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Senha,
			&usuario.CriadoEm,
		)
		if erro != nil {
			return nil, erro
		}

		// Adiciona o usuário ao slice de usuários
		usuarios = append(usuarios, usuario)
	}

	// Verifica se ocorreu algum erro durante a iteração das linhas
	if erro = query.Err(); erro != nil {
		return nil, erro
	}

	return usuarios, nil
}

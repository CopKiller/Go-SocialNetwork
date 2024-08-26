package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

func (repositorio Usuarios) BuscarUsuarios(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	query, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM usuarios WHERE Nome LIKE ? OR Nick LIKE ?",
		nomeOuNick, nomeOuNick)
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

func (repositorio Usuarios) BuscarUsuarioID(ID uint64) (modelos.Usuario, error) {
	query, erro := repositorio.db.Query(
		"SELECT ID, NOME, NICK, EMAIL, CRIADOEM FROM usuarios WHERE ID = ?",
		ID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer query.Close()

	var usuario modelos.Usuario
	// Itera sobre as linhas retornadas pela query
	for query.Next() {
		// Escaneia cada coluna do resultado para os campos correspondentes do struct Usuario
		erro = query.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio Usuarios) AtualizarUsuario(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET NOME = ?, NICK = ?, EMAIL = ? WHERE ID = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}
	return nil
}

func (repositorio Usuarios) DeletarUsuario(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE ID = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT ID, SENHA FROM usuarios WHERE EMAIL = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	for linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

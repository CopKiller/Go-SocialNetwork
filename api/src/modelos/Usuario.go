package modelos

import (
	"api/src/seguranca"
	"github.com/badoux/checkmail"
	"gopkg.in/validator.v2"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty" validate:"nonzero,max=50"`
	Nick     string    `json:"nick,omitempty" validate:"min=3,max=50,regexp=^[a-zA-Z0-9_]*$"`
	Email    string    `json:"email,omitempty" validate:"min=3,max=50,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Senha    string    `json:"senha,omitempty" validate:"min=3,max=100"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) CheckDataRules(etapa string) error {

	if etapa != "cadastro" {
		usuario.Senha = "***"
	}

	err := validator.Validate(usuario)
	if err != nil {
		return err
	}

	err = checkmail.ValidateFormat(usuario.Email)
	if err != nil {
		return err
	}

	if err = usuario.formatar(etapa); err != nil {
		return err
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}

	return nil
}

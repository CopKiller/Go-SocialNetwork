package modelos

import (
	"github.com/badoux/checkmail"
	"gopkg.in/validator.v2"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty" validate:"nonzero,max=50"`
	Nick     string    `json:"nick,omitempty" validate:"min=3,max=50,regexp=^[a-zA-Z]*$"`
	Email    string    `json:"email,omitempty" validate:"min=3,max=50,regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Senha    string    `json:"senha,omitempty" validate:"min=3,max=20"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

func (usuario *Usuario) CheckDataRules() error {
	err := validator.Validate(usuario)
	if err != nil {
		return err
	}

	err = checkmail.ValidateFormat(usuario.Email)
	if err != nil {
		return err
	}

	return nil
}

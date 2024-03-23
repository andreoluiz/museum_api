package internal

import (
	"github.com/google/uuid"
)

type Post struct {
	ID              uuid.UUID `json:"-"`
	Nome            string    `json:"nome"`
	Data_nascimento string    `json:"data_nascimento"`
	Cpf             string    `json:"cpf"`
	Senha           string    `json:"senha"`
}

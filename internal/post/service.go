package post

import (
	"errors"

	"github.com/VirtualArtExplore/api/internal"
	"github.com/google/uuid"
)

var ErrPostNameEmpty = errors.New("post name is empty")
var ErrPostPasswordEmpty = errors.New("post password ist limit")
var ErrPostNotFound = errors.New("post id not found")

type Service struct {
	Repository Repository
}

func (p Service) Create(post internal.Post) error {
	if post.Nome == "" {
		return ErrPostNameEmpty
	}

	if post.Senha == "" {
		return ErrPostPasswordEmpty
	}

	return p.Repository.Insert(post)
}

func (p Service) Delete(id uuid.UUID) error {
	err := p.Repository.Delete(id)
	if err != nil {
		if errors.Is(err, ErrPostNotFound) {
			return ErrPostNotFound
		}
		return err
	}

	return nil
}

package post

import (
	"errors"
	"unicode/utf8"

	"github.com/VirtualArtExplore/api/internal"
	"github.com/google/uuid"
)

var ErrPostBodyEmpty = errors.New("post body is empty")
var ErrPostBodyExceedLimit = errors.New("post body exceeds limit")
var ErrPostNotFound = errors.New("post id not found")

type Service struct {
	Repository Repository
}

func (p Service) Create(post internal.Post) error {
	if post.Body == "" {
		return ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 148 {
		return ErrPostBodyExceedLimit
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

package usecase

import (
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IUseCase interface {
	CreateTinyURL(url entity.OriginalURL) (entity.TinyURL, error)
	GetOriginalURLByTinyURL(tinyURL entity.TinyURL) (entity.OriginalURL, error)
}

var (
	ErrNotFound = errors.New("not found")
)

package usecase

import (
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IUseCase interface {
	CreateTinyURL(url entity.OriginURL) (entity.TinyURL, error)
	GetTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error)
}

var (
	ErrNotFound = errors.New("not found")
)

package usecase

import (
	"context"
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IUseCase interface {
	CreateTinyURL(ctx *context.Context, url entity.OriginalURL) (entity.TinyURL, error)
	GetOriginalURLByTinyURL(ctx *context.Context, tinyURL entity.TinyURL) (entity.OriginalURL, error)
}

var (
	ErrNotFound = errors.New("not found")
)

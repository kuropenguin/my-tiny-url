package repository

import (
	"context"
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IRepository interface {
	Save(ctx *context.Context, url entity.OriginalURL, tinyURL entity.TinyURL) error
	FindOriginalURLByTinyURL(ctx *context.Context, tinyURL entity.TinyURL) (entity.OriginalURL, error)
	FindTinyURLByURL(ctx *context.Context, URL entity.OriginalURL) (entity.TinyURL, error)
}

type ICacheRepository interface {
	Set(string, string) error
	Get(string) (string, error)
}

var (
	ErrNotFound      = errors.New("not found")
	ErrCacheNotFound = errors.New("cache not found")
)

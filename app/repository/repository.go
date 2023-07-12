package repository

import (
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IRepository interface {
	Save(url entity.OriginalURL, tinyURL entity.TinyURL) error
	FindOriginalURLbyTinyURL(tinyURL entity.TinyURL) (entity.OriginalURL, error)
	FindTinyURLbyURL(URL entity.OriginalURL) (entity.TinyURL, error)
}

type ICacheRepository interface {
	Set(string, string) error
	Get(string) (string, error)
}

var (
	ErrNotFound      = errors.New("not found")
	ErrCacheNotFound = errors.New("cache not found")
)

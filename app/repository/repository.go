package repository

import (
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IRepository interface {
	Save(url entity.OriginURL, tinyURL entity.TinyURL) error
	FindOriginURLbyTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error)
	FindTinyURLbyURL(URL entity.OriginURL) (entity.TinyURL, error)
}

type ICacheRepository interface {
	Save(url entity.OriginURL, tinyURL entity.TinyURL) error
	GetOriginalURLByTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error)
	GetTinyURLByOriginalURL(tinyURL entity.OriginURL) (entity.TinyURL, error)
}

var (
	ErrNotFound = errors.New("not found")
)

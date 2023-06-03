package repository

import (
	"errors"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

type IRepository interface {
	Create(url entity.OriginURL, tinyURL entity.TinyURL) error
	FindbyTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error)
	FindbyURL(URL entity.OriginURL) (entity.TinyURL, error)
}

var (
	ErrNotFound = errors.New("not found")
)

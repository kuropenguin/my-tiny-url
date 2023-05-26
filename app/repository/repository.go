package repository

import "errors"

type IRepository interface {
	Create(url string, tinyURL string) error
	FindbyTinyURL(tinyURL string) (string, error)
	FindbyURL(URL string) (string, error)
}

var (
	ErrNotFound = errors.New("not found")
)

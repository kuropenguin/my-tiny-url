package usecase

import "errors"

type IUseCase interface {
	CreateTinyURL(url string) (string, error)
	GetTinyURL(tinyURL string) (string, error)
}

var (
	ErrNotFound = errors.New("not found")
)

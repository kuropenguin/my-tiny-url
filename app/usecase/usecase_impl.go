package usecase

import (
	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/repository"
)

type UsecaseImpl struct {
	repository repository.IRepository
}

func NewUsecaseImpl(repository repository.IRepository) *UsecaseImpl {
	return &UsecaseImpl{repository: repository}
}

func (u *UsecaseImpl) CreateTinyURL(url entity.OriginURL) (entity.TinyURL, error) {
	tinyURL, err := u.repository.FindbyURL(url)
	// 既にあるならそれを返す
	if err == nil {
		return tinyURL, nil
	}
	for {
		tinyURL = entity.GenerateTinyURL()
		if _, err := u.repository.FindbyTinyURL(tinyURL); err != nil {
			if err != repository.ErrNotFound {
				return "", err
			}
			err = u.repository.Create(url, tinyURL)
			if err != nil {
				return "", err
			}
			break
		}
	}
	return tinyURL, nil
}

func (u *UsecaseImpl) GetByTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error) {
	url, err := u.repository.FindbyTinyURL(tinyURL)
	if err != nil {
		return "", ErrNotFound
	}
	return url, nil
}

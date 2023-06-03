package usecase

import "github.com/kuropenguin/my-tiny-url/app/repository"

type UsecaseImpl struct {
	repository repository.IRepository
}

func NewUsecaseImpl(repository repository.IRepository) *UsecaseImpl {
	return &UsecaseImpl{repository: repository}
}

func (u *UsecaseImpl) CreateTinyURL(url string) (string, error) {
	tinyURL, err := u.repository.FindbyURL(url)
	// 既にあるならそれを返す
	if err == nil {
		return tinyURL, nil
	}
	err = u.repository.Create(url, tinyURL)
	if err != nil {
		return "", err
	}
	return tinyURL, nil
}

func (u *UsecaseImpl) GetTinyURL(tinyURL string) (string, error) {
	url, err := u.repository.FindbyTinyURL(tinyURL)
	if err != nil {
		return "", ErrNotFound
	}
	return url, nil
}

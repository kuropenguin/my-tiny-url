package usecase

import (
	"log"

	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/repository"
)

type UsecaseImpl struct {
	repository repository.IRepository
	cache      repository.ICacheRepository
}

func NewUsecaseImpl(repository repository.IRepository, cache repository.ICacheRepository) *UsecaseImpl {
	return &UsecaseImpl{
		repository: repository,
		cache:      cache,
	}
}

func (u *UsecaseImpl) CreateTinyURL(url entity.OriginalURL) (entity.TinyURL, error) {

	// check cache
	cachedTinyURL, err := u.cache.Get(string(url))
	if err != nil && err != repository.ErrCacheNotFound {
		return "", err
	}
	if cachedTinyURL != "" {
		return entity.TinyURL(cachedTinyURL), nil
	}

	tinyURL, err := u.repository.FindTinyURLbyURL(url)
	// 既にあるならそれを返す
	if err == nil {
		return tinyURL, nil
	}

	for {
		tinyURL = entity.GenerateTinyURL()
		if _, err := u.repository.FindOriginalURLbyTinyURL(tinyURL); err != nil {
			if err != repository.ErrNotFound {
				return "", err
			}
			err = u.repository.Save(url, tinyURL)
			if err != nil {
				return "", err
			}
			break
		}
	}
	u.cache.Set(string(url), string(tinyURL))
	u.cache.Set(string(tinyURL), string(url))
	return tinyURL, nil
}

func (u *UsecaseImpl) GetOriginalURLByTinyURL(tinyURL entity.TinyURL) (entity.OriginalURL, error) {
	cachedOriginalURL, err := u.cache.Get(string(tinyURL))
	if err != nil && err != repository.ErrCacheNotFound {
		return "", err
	}
	if cachedOriginalURL != "" {
		return entity.OriginalURL(cachedOriginalURL), nil
	}

	url, err := u.repository.FindOriginalURLbyTinyURL(tinyURL)
	if err == repository.ErrNotFound {
		return "", ErrNotFound
	}
	if err != nil {
		return "", err
	}

	err = u.cache.Set(string(tinyURL), string(url))
	if err != nil {
		log.Println("save err", err)
	}
	return url, nil
}

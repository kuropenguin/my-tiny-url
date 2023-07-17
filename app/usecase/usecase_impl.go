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

	// 既存DBチェック
	tinyURL, err := u.repository.FindTinyURLByURL(url)
	// 既にあるならそれを返す
	if err == nil {
		return tinyURL, nil
	}
	if err != repository.ErrNotFound {
		return "", err
	}

	for {
		// 作成
		tinyURL = entity.GenerateTinyURL()
		// 重複チェック
		if _, err := u.repository.FindOriginalURLByTinyURL(tinyURL); err != nil {
			if err != repository.ErrNotFound {
				return "", err
			}
			// 重複していなければ保存(ここに来るのは ErrNotFound のみ)
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

	url, err := u.repository.FindOriginalURLByTinyURL(tinyURL)
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

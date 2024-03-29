package repository

import (
	"context"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

func NewMapRepository() IRepository {
	return &MapRepository{
		// TODO 永続化層に移す
		//参考にする https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
		URLStorage: make(map[entity.TinyURL]entity.OriginalURL),
	}
}

type MapRepository struct {
	URLStorage map[entity.TinyURL]entity.OriginalURL
}

func (m *MapRepository) Save(ctx *context.Context, url entity.OriginalURL, tinyURL entity.TinyURL) error {
	m.URLStorage[tinyURL] = url
	return nil
}

func (m *MapRepository) FindOriginalURLByTinyURL(ctx *context.Context, tinyURL entity.TinyURL) (entity.OriginalURL, error) {
	if _, ok := m.URLStorage[tinyURL]; ok {
		return m.URLStorage[tinyURL], nil
	}
	return "", ErrNotFound
}

func (m *MapRepository) FindTinyURLByURL(ctx *context.Context, url entity.OriginalURL) (entity.TinyURL, error) {
	for tinyURL, originalURL := range m.URLStorage {
		if originalURL == url {
			return tinyURL, nil
		}
	}
	return "", ErrNotFound
}

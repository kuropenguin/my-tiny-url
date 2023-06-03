package repository

import (
	"github.com/kuropenguin/my-tiny-url/app/entity"
)

func NewMapRepository() IRepository {
	return &MapRepository{
		// TODO 永続化層に移す
		//参考にする https://qiita.com/hirotakan/items/698c1f5773a3cca6193e
		URLStorage: make(map[entity.TinyURL]entity.OriginURL),
	}
}

type MapRepository struct {
	URLStorage map[entity.TinyURL]entity.OriginURL
}

func (m *MapRepository) Save(url entity.OriginURL, tinyURL entity.TinyURL) error {
	m.URLStorage[tinyURL] = url
	return nil
}

func (m *MapRepository) FindOriginURLbyTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error) {
	if _, ok := m.URLStorage[tinyURL]; ok {
		return m.URLStorage[tinyURL], nil
	}
	return "", ErrNotFound
}

func (m *MapRepository) FindTinyURLbyURL(url entity.OriginURL) (entity.TinyURL, error) {
	for tinyURL, originURL := range m.URLStorage {
		if originURL == url {
			return tinyURL, nil
		}
	}
	return "", ErrNotFound
}

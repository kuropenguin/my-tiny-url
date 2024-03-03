package repository

import (
	"context"
	"database/sql"

	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/sqlc/queries"
)

func NewMysqlRepository(q *queries.Queries) IRepository {
	return &MysqlRepository{
		queries: q,
	}
}

type MysqlRepository struct {
	queries *queries.Queries
}

func (m *MysqlRepository) Save(url entity.OriginalURL, tinyURL entity.TinyURL) error {
	if _, err := m.queries.CreateURLs(context.Background(), queries.CreateURLsParams{
		OriginalUrl: string(url),
		TinyUrl:     string(tinyURL),
	}); err != nil {
		return err
	}
	return nil
}

func (m *MysqlRepository) FindOriginalURLByTinyURL(tinyURL entity.TinyURL) (entity.OriginalURL, error) {
	rawOriginalURL, err := m.queries.GetOriginalURLByTinyURL(context.Background(), string(tinyURL))
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return entity.OriginalURL(rawOriginalURL), nil
}

func (m *MysqlRepository) FindTinyURLByURL(url entity.OriginalURL) (entity.TinyURL, error) {
	rawTinyURL, err := m.queries.GetTinyURLByOriginalURL(context.Background(), string(url))
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return entity.TinyURL(rawTinyURL), nil
}

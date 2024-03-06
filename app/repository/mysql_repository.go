package repository

import (
	"context"
	"database/sql"

	"github.com/kuropenguin/my-tiny-url/app/config"
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

func getTx(ctx *context.Context) *queries.Queries {
	return (*ctx).Value(config.TxKey).(*queries.Queries)
}

func (m *MysqlRepository) Save(ctx *context.Context, url entity.OriginalURL, tinyURL entity.TinyURL) error {
	tx := getTx(ctx)
	if _, err := tx.CreateURLs(context.Background(), queries.CreateURLsParams{
		OriginalUrl: string(url),
		TinyUrl:     string(tinyURL),
	}); err != nil {
		return err
	}
	return nil
}

func (m *MysqlRepository) FindOriginalURLByTinyURL(ctx *context.Context, tinyURL entity.TinyURL) (entity.OriginalURL, error) {
	rawOriginalURL, err := m.queries.GetOriginalURLByTinyURL(context.Background(), string(tinyURL))
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return entity.OriginalURL(rawOriginalURL), nil
}

func (m *MysqlRepository) FindTinyURLByURL(ctx *context.Context, url entity.OriginalURL) (entity.TinyURL, error) {
	rawTinyURL, err := m.queries.GetTinyURLByOriginalURL(context.Background(), string(url))
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return entity.TinyURL(rawTinyURL), nil
}

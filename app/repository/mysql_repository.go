package repository

import (
	"database/sql"

	"github.com/kuropenguin/my-tiny-url/app/entity"
)

func NewMysqlRepository(db *sql.DB) IRepository {
	return &MysqlRepository{
		URLStorage: db,
	}
}

type MysqlRepository struct {
	URLStorage *sql.DB
}

func (m *MysqlRepository) Save(url entity.OriginalURL, tinyURL entity.TinyURL) error {
	stmt, err := m.URLStorage.Prepare("INSERT INTO urls (original_url, tiny_url) VALUES (?, ?)")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(url, tinyURL)
	if err != nil {
		return err
	}
	if _, err := result.LastInsertId(); err != nil {
		return err
	}
	return nil
}

func (m *MysqlRepository) FindOriginalURLbyTinyURL(tinyURL entity.TinyURL) (entity.OriginalURL, error) {
	row := m.URLStorage.QueryRow("SELECT original_url FROM urls WHERE tiny_url = ?", tinyURL)
	if row.Err() != nil {
		return "", row.Err()
	}
	var originalURL entity.OriginalURL
	if err := row.Scan(&originalURL); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}

	// sqlc example
	// queries := tutorial.New(m.URLStorage)
	// urls, _ := queries.GetTinyURL(nil, string(originalURL))

	return originalURL, nil
}

func (m *MysqlRepository) FindTinyURLbyURL(url entity.OriginalURL) (entity.TinyURL, error) {
	row := m.URLStorage.QueryRow("SELECT tiny_url FROM urls WHERE original_url = ?", url)
	if row.Err() != nil {
		return "", row.Err()
	}
	var tinyURL entity.TinyURL
	if err := row.Scan(&tinyURL); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return tinyURL, nil
}

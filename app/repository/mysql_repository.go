package repository

import (
	"database/sql"
	"fmt"

	"github.com/kuropenguin/my-tiny-url/app/entity"
	"github.com/kuropenguin/my-tiny-url/app/tutorial"
)

func NewMysqlRepository(db *sql.DB) IRepository {
	return &MysqlRepository{
		URLStorage: db,
	}
}

type MysqlRepository struct {
	URLStorage *sql.DB
}

func (m *MysqlRepository) Save(url entity.OriginURL, tinyURL entity.TinyURL) error {
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
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return nil
}

func (m *MysqlRepository) FindOriginURLbyTinyURL(tinyURL entity.TinyURL) (entity.OriginURL, error) {
	row := m.URLStorage.QueryRow("SELECT original_url FROM urls WHERE tiny_url = ?", tinyURL)
	if row.Err() != nil {
		return "", row.Err()
	}
	var originalURL entity.OriginURL
	if err := row.Scan(&originalURL); err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}

	queries := tutorial.New(m.URLStorage)
	urls, _ := queries.GetTinyURL(nil, string(originalURL))
	fmt.Println(urls)

	return originalURL, nil
}

func (m *MysqlRepository) FindTinyURLbyURL(url entity.OriginURL) (entity.TinyURL, error) {
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

package repository

import (
	"database/sql"
	"fmt"
	"log"

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

func (m *MysqlRepository) Save(url entity.OriginURL, tinyURL entity.TinyURL) error {
	stmt, err := m.URLStorage.Prepare("INSERT INTO url_storage (tiny_url, origin_url) VALUES (?, ?)")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(tinyURL, url)
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
	row := m.URLStorage.QueryRow("SELECT origin_url FROM url_storage WHERE tiny_url = ?", tinyURL)
	if row.Err() != nil {
		log.Println(row.Err())
		return "", row.Err()
	}
	var originalURL entity.OriginURL
	if err := row.Scan(&originalURL); err != nil {
		log.Println(err)
		return "", err
	}
	return originalURL, nil
}

func (m *MysqlRepository) FindTinyURLbyURL(url entity.OriginURL) (entity.TinyURL, error) {
	row := m.URLStorage.QueryRow("SELECT tiny_url FROM url_storage WHERE original_url = ?", url)
	if row.Err() != nil {
		log.Println(row.Err())
		return "", row.Err()
	}
	var tinyURL entity.TinyURL
	if err := row.Scan(&tinyURL); err != nil {
		log.Println(err)
		return "", err
	}
	return tinyURL, nil
}

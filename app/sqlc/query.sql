-- name: GetTinyURLByOriginalURL :one

SELECT tiny_url FROM urls WHERE original_url = ?;

-- name: GetOriginalURLByTinyURL :one

SELECT original_url FROM urls WHERE tiny_url = ?;

-- name: CreateURLs :execresult

INSERT INTO urls (original_url, tiny_url) VALUES (?, ?);

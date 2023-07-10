-- name: GetTinyURL :one

SELECT * FROM urls WHERE original_url = ? LIMIT 1;

-- name: ListOriginalURL :many

SELECT * FROM urls ORDER BY original_url;

-- name: CreateAuthor :execresult

INSERT INTO urls ( original_url, tiny_url ) VALUES ( ?, ? );

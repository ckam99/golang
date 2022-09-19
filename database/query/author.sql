-- name: CreateAuthor :one
INSERT INTO authors (
    name, bio
) VALUES(
    $1,$2
) RETURNING *;


-- name: GetAuthor :one
SELECT * FROM authors WHERE id=$1 LIMIT 1;

-- name: GetAllAuthors :many
SELECT * FROM authors ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAuthor :one
UPDATE authors SET 
name = COALESCE(sqlc.narg('name'), name),
bio = COALESCE(sqlc.narg('bio'), bio),
updated_at = now() 
WHERE id = sqlc.arg('id') RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id=$1;

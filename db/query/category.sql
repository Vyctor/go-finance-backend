-- name: CreateCategory :one
INSERT INTO categories (user_id, title, type, description)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetCategoryByUserId :one
SELECT *
FROM categories
WHERE user_id = $1
LIMIT 1;

-- name: GetCategoryById :one
SELECT *
FROM categories
WHERE id = $1
LIMIT 1;

-- name: GetCategories :many
SELECT *
FROM categories
WHERE user_id = $1
  and type = $2
  and title like $3
  and description like $4;

-- name: UpdateCategory :one
UPDATE categories
SET title       = $2,
    type        = $3,
    description = $4,
    updated_at  = now()
WHERE id = $1
RETURNING id;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;


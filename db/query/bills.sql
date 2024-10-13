-- name: CreateBill :one
INSERT INTO bills (user_id, category_id, title, type, description, amount, date)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetBill :one
SELECT *
FROM bills
WHERE id = $1
LIMIT 1;

-- name: GetBills :many
SELECT b.*, c.title as category_title
FROM bills b
         LEFT JOIN categories c ON c.id = b.category_id
WHERE b.user_id = $1
  and b.type = $2
  and b.title like $3
  and b.description like $4
  and b.category_id = $5
  and b.date = $6
ORDER BY date DESC;

-- name: GetBillsGraph :one
SELECT COUNT(*)
FROM bills
where user_id = $1
  and type = $2;

-- name: GetBillsReports :one
SELECT SUM(value) AS sum_value
FROM bills
where user_id = $1
  and type = $2;

-- name: UpdateBill :one

UPDATE bills
SET title       = $2,
    type        = $3,
    description = $4,
    amount      = $5,
    date        = $6
WHERE id = $1
RETURNING id;

-- name: DeleteBill :exec
DELETE
FROM bills
WHERE id = $1;


-- name: CreateTask :one
INSERT INTO tasks (id, status,type,expiry_date,created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: CreateCatatan :one
INSERT INTO catatan
(id_akun, jumlah)
VALUES
($1, $2)
RETURNING *;

-- name: ListCatatan :many
SELECT * FROM catatan
WHERE id_akun = $1
LIMIT $1
OFFSET $2;
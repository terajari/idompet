-- name: CreateAkun :one
INSERT INTO akun
(nama, saldo)
VALUES
($1, $2)
RETURNING *;

-- name: GetAkun :one
SELECT * FROM akun
WHERE id = $1;

-- name: GetAkunForUpdate :one
SELECT * FROM akun
WHERE id = $1
FOR NO KEY UPDATE;

-- name: ListAkun :many
SELECT * FROM akun
LIMIT $1
OFFSET $2;

-- name: UpdateAkun :one
UPDATE akun
SET saldo = $2
WHERE id = $1
RETURNING *;

-- name: ChangeSaldo :one
UPDATE akun
SET saldo = saldo + sqlc.arg(jumlah)
WHERE id = $1
RETURNING *;

-- name: DeleteAkun :exec
DELETE FROM akun
WHERE id = $1;
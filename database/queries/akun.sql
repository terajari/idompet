-- name: CreateAkun :one
INSERT INTO akun
(nama, saldo)
VALUES
($1, $2)
RETURNING *;

-- name: GetAkun :one
SELECT * FROM akun
WHERE id = $1;

-- name: ListAkun :many
SELECT * FROM akun
LIMIT $1
OFFSET $2;

-- name: UpdateAkun :exec
UPDATE akun
SET saldo = $2
WHERE id = $1;

-- name: DeleteAkun :exec
DELETE FROM akun
WHERE id = $1;
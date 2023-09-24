-- name: CreateTransfer :one
INSERT INTO transfer
(pengirim, penerima, jumlah)
VALUES
($1, $2, $3)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM catatan
WHERE id = $1;
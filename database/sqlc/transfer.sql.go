// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: transfer.sql

package database

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfer
(pengirim, penerima, jumlah)
VALUES
($1, $2, $3)
RETURNING id, pengirim, penerima, jumlah, dibuat_pada
`

type CreateTransferParams struct {
	Pengirim int64 `json:"pengirim"`
	Penerima int64 `json:"penerima"`
	Jumlah   int64 `json:"jumlah"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.Pengirim, arg.Penerima, arg.Jumlah)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.Pengirim,
		&i.Penerima,
		&i.Jumlah,
		&i.DibuatPada,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, pengirim, penerima, jumlah, dibuat_pada FROM transfer
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.Pengirim,
		&i.Penerima,
		&i.Jumlah,
		&i.DibuatPada,
	)
	return i, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: catatan.sql

package database

import (
	"context"
)

const createCatatan = `-- name: CreateCatatan :one
INSERT INTO catatan
(id_akun, jumlah)
VALUES
($1, $2)
RETURNING id, id_akun, jumlah, dibuat_pada
`

type CreateCatatanParams struct {
	IDAkun int64 `json:"id_akun"`
	Jumlah int64 `json:"jumlah"`
}

func (q *Queries) CreateCatatan(ctx context.Context, arg CreateCatatanParams) (Catatan, error) {
	row := q.db.QueryRowContext(ctx, createCatatan, arg.IDAkun, arg.Jumlah)
	var i Catatan
	err := row.Scan(
		&i.ID,
		&i.IDAkun,
		&i.Jumlah,
		&i.DibuatPada,
	)
	return i, err
}

const listCatatan = `-- name: ListCatatan :many
SELECT id, id_akun, jumlah, dibuat_pada FROM catatan
WHERE id_akun = $1
LIMIT $1
OFFSET $2
`

type ListCatatanParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCatatan(ctx context.Context, arg ListCatatanParams) ([]Catatan, error) {
	rows, err := q.db.QueryContext(ctx, listCatatan, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Catatan{}
	for rows.Next() {
		var i Catatan
		if err := rows.Scan(
			&i.ID,
			&i.IDAkun,
			&i.Jumlah,
			&i.DibuatPada,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

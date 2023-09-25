package database

import (
	"context"
	"database/sql"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(sql *sql.DB) *Store {
	return &Store{
		db:      sql,
		Queries: New(sql),
	}
}

// argumen PerformTransfer
type TransferTxParams struct {
	IDPengirim int64 `json:"id_pengirim"`
	IDPenerima int64 `json:"id_penerima"`
	Jumlah     int64 `json:"jumlah"`
}

// hasil PerformTransfer
type TransferTxResult struct {
	Transfer        Transfer `json:"transfer"`
	Pengirim        Akun     `json:"pengirim"`
	Penerima        Akun     `json:"penerima"`
	CatatanPengirim Catatan  `json:"catatan_pengirim"`
	CatatanPenerima Catatan  `json:"catatan_penerima"`
}

// Transfer saldo dari pengirim ke penerima
// lalu mencatat mutasi
func (store *Store) PerformTransfer(ctx context.Context, arg TransferTxParams) (*TransferTxResult, error) {
	var result TransferTxResult
	tx, err := store.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	qtx := store.WithTx(tx)

	result.Transfer, err = qtx.CreateTransfer(ctx, CreateTransferParams{
		Pengirim: arg.IDPengirim,
		Penerima: arg.IDPenerima,
		Jumlah:   arg.Jumlah,
	})
	if err != nil {
		return nil, err
	}

	result.CatatanPengirim, err = qtx.CreateCatatan(ctx, CreateCatatanParams{
		IDAkun: arg.IDPengirim,
		Jumlah: -arg.Jumlah,
	})
	if err != nil {
		return nil, err
	}

	result.CatatanPenerima, err = qtx.CreateCatatan(ctx, CreateCatatanParams{
		IDAkun: arg.IDPenerima,
		Jumlah: arg.Jumlah,
	})

	if err != nil {
		return nil, err
	}

	// GetAkun ->UpdateAkun

	result.Pengirim, err = qtx.ChangeSaldo(ctx, ChangeSaldoParams{
		ID:     arg.IDPengirim,
		Jumlah: -arg.Jumlah,
	})
	if err != nil {
		return nil, err
	}

	result.Penerima, err = qtx.ChangeSaldo(ctx, ChangeSaldoParams{
		ID:     arg.IDPenerima,
		Jumlah: arg.Jumlah,
	})
	if err != nil {
		return nil, err
	}

	return &result, tx.Commit()
}

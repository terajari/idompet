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
	if arg.IDPengirim < arg.IDPenerima {
		result.Pengirim, result.Penerima, err = AddMoney(ctx, store, arg.IDPengirim, -arg.Jumlah, arg.IDPenerima, arg.Jumlah)
		if err != nil {
			return nil, err
		}
	} else {
		result.Penerima, result.Pengirim, err = AddMoney(ctx, store, arg.IDPenerima, arg.Jumlah, arg.IDPengirim, -arg.Jumlah)
		if err != nil {
			return nil, err
		}
	}

	return &result, tx.Commit()
}

func AddMoney(
	ctx context.Context,
	store *Store,
	acc1,
	jumlah1,
	acc2,
	jumlah2 int64,
) (account1 Akun, account2 Akun, err error) {
	account1, err = store.ChangeSaldo(ctx, ChangeSaldoParams{
		ID:     acc1,
		Jumlah: jumlah1,
	})
	if err != nil {
		return
	}

	account2, err = store.ChangeSaldo(ctx, ChangeSaldoParams{
		ID:     acc2,
		Jumlah: jumlah2,
	})
	return
}

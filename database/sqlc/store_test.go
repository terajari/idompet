package database

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerformTransfer(t *testing.T) {
	pengirim := createRandomAkun(t)
	penerima := createRandomAkun(t)

	fmt.Println("before:", pengirim.Saldo, penerima.Saldo)
	errs := make(chan error)
	results := make(chan TransferTxResult)

	n := int64(3)
	jumlah := int64(500)

	for i := 0; int64(i) < n; i++ {
		go func() {
			result, err := testQueries.PerformTransfer(context.Background(), TransferTxParams{
				IDPengirim: pengirim.ID,
				IDPenerima: penerima.ID,
				Jumlah:     jumlah,
			})

			errs <- err
			results <- *result
		}()
	}

	existed := make(map[int]bool)
	// check result
	for i := 0; int64(i) < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		tranfer := result.Transfer
		require.NotEmpty(t, tranfer)
		require.Equal(t, pengirim.ID, tranfer.Pengirim)
		require.Equal(t, penerima.ID, tranfer.Penerima)
		require.Equal(t, jumlah, tranfer.Jumlah)
		require.NotZero(t, tranfer.ID)
		require.NotZero(t, tranfer.DibuatPada)

		_, err = testQueries.GetTransfer(context.Background(), tranfer.ID)
		require.NoError(t, err)

		// check catatan pengirim
		catatanPengirim := result.CatatanPengirim
		require.NotEmpty(t, catatanPengirim)
		require.Equal(t, pengirim.ID, catatanPengirim.IDAkun)
		require.Equal(t, -jumlah, catatanPengirim.Jumlah)
		require.NotZero(t, catatanPengirim.ID)
		require.NotZero(t, catatanPengirim.DibuatPada)

		_, err = testQueries.GetCatatan(context.Background(), catatanPengirim.ID)
		require.NoError(t, err)

		// check catatan penerima
		catatanPenerima := result.CatatanPenerima
		require.NotEmpty(t, catatanPenerima)
		require.Equal(t, penerima.ID, catatanPenerima.IDAkun)
		require.Equal(t, jumlah, catatanPenerima.Jumlah)
		require.NotZero(t, catatanPenerima.ID)
		require.NotZero(t, catatanPenerima.DibuatPada)

		_, err = testQueries.GetCatatan(context.Background(), catatanPenerima.ID)
		require.NoError(t, err)

		// check akun
		akunPengirim := result.Pengirim
		require.NotEmpty(t, akunPengirim)
		require.Equal(t, pengirim.ID, akunPengirim.ID)

		akunPenerima := result.Penerima
		require.NotEmpty(t, akunPenerima)
		require.Equal(t, penerima.ID, akunPenerima.ID)

		// check saldo
		selisihSaldoPengirim := pengirim.Saldo - akunPengirim.Saldo
		selisihSaldoPenerima := akunPenerima.Saldo - penerima.Saldo
		require.Equal(t, selisihSaldoPengirim, selisihSaldoPenerima)
		require.True(t, selisihSaldoPengirim > 0)
		require.True(t, selisihSaldoPengirim%jumlah == 0)

		k := int(selisihSaldoPengirim / jumlah)
		require.True(t, k >= 1 && k <= int(n))
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check updated balance
	updatedAkunPengirim, err := testQueries.GetAkun(context.Background(), pengirim.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAkunPengirim)

	updatedAkunPenerima, err := testQueries.GetAkun(context.Background(), penerima.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAkunPenerima)

	fmt.Println("after:", updatedAkunPengirim.Saldo, updatedAkunPenerima.Saldo)
	require.Equal(t, pengirim.Saldo-n*jumlah, updatedAkunPengirim.Saldo)
	require.Equal(t, penerima.Saldo+n*jumlah, updatedAkunPenerima.Saldo)
}

func TestPerformTransferDeadlock(t *testing.T) {
	pengirim := createRandomAkun(t)
	penerima := createRandomAkun(t)

	fmt.Println("before:", pengirim.Saldo, penerima.Saldo)
	errs := make(chan error)

	n := int64(10)
	jumlah := int64(500)

	for i := 0; int64(i) < n; i++ {
		pengirimID := pengirim.ID
		penerimaID := penerima.ID

		if i%2 == 1 {
			pengirimID = penerima.ID
			penerimaID = pengirim.ID
		}

		go func() {
			_, err := testQueries.PerformTransfer(context.Background(), TransferTxParams{
				IDPengirim: pengirimID,
				IDPenerima: penerimaID,
				Jumlah:     jumlah,
			})

			errs <- err
		}()
	}

	// check result
	for i := 0; int64(i) < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check updated balance
	updatedAkunPengirim, err := testQueries.GetAkun(context.Background(), pengirim.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAkunPengirim)

	updatedAkunPenerima, err := testQueries.GetAkun(context.Background(), penerima.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAkunPenerima)

	fmt.Println("after:", updatedAkunPengirim.Saldo, updatedAkunPenerima.Saldo)
	require.Equal(t, pengirim.Saldo, updatedAkunPengirim.Saldo)
	require.Equal(t, penerima.Saldo, updatedAkunPenerima.Saldo)
}

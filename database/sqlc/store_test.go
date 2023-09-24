package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPerformTransfer(t *testing.T) {
	pengirim := createRandomAkun(t)
	penerima := createRandomAkun(t)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	n := 5
	jumlah := int64(10)

	for i := 0; i < n; i++ {
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

	for i := 0; i < n; i++ {
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
	}
}

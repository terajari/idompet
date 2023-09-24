package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/terajari/idompet/util"
)

func createRandomAkun(t *testing.T) Akun {
	arg := CreateAkunParams{
		Nama:  util.RandomName(),
		Saldo: util.RandomSaldo(),
	}

	akun, err := testQueries.CreateAkun(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, akun)

	require.Equal(t, akun.Nama, arg.Nama)
	require.Equal(t, akun.Saldo, arg.Saldo)

	require.NotZero(t, akun.ID)
	require.NotZero(t, akun.DibuatPada)

	return akun
}

func TestCreateAkun(t *testing.T) {
	createRandomAkun(t)
}

func TestGetAkun(t *testing.T) {
	akun := createRandomAkun(t)

	getAkun, err := testQueries.GetAkun(context.Background(), akun.ID)
	require.NoError(t, err)

	require.Equal(t, getAkun.ID, akun.ID)
	require.Equal(t, getAkun.Nama, akun.Nama)
	require.Equal(t, getAkun.Saldo, akun.Saldo)
	require.WithinDuration(t, getAkun.DibuatPada, akun.DibuatPada, time.Millisecond)
}

func TestListAkun(t *testing.T) {
	arg := ListAkunParams{
		Limit:  5,
		Offset: 5,
	}

	for i := 0; i < int(arg.Limit+arg.Offset); i++ {
		createRandomAkun(t)
	}

	listAkun, err := testQueries.ListAkun(context.Background(), arg)
	require.NoError(t, err)

	for _, akun := range listAkun {
		require.NotEmpty(t, akun)
	}
}

func TestUpdateAkun(t *testing.T) {
	akun := createRandomAkun(t)
	arg := UpdateAkunParams{
		ID:    akun.ID,
		Saldo: util.RandomSaldo(),
	}

	updatedAkun, err := testQueries.UpdateAkun(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, updatedAkun.Nama, akun.Nama)
	require.Equal(t, updatedAkun.Saldo, arg.Saldo)
	require.Equal(t, updatedAkun.ID, akun.ID)
	require.Equal(t, updatedAkun.DibuatPada, akun.DibuatPada)

}

func TestDeleteAkun(t *testing.T) {
	akun := createRandomAkun(t)

	err := testQueries.DeleteAkun(context.Background(), akun.ID)
	require.NoError(t, err)

	getAkun, err := testQueries.GetAkun(context.Background(), akun.ID)
	require.Error(t, err)
	require.Empty(t, getAkun)
}

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdatabase "github.com/terajari/idompet/database/mock"
	database "github.com/terajari/idompet/database/sqlc"
	"github.com/terajari/idompet/util"
)

func TestGetAkunAPI(t *testing.T) {
	akun := randomAkun()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdatabase.NewMockStore(ctrl)

	mockStore.EXPECT().
		GetAkun(gomock.Any(), gomock.Eq(akun.ID)).
		Times(1).
		Return(akun, nil)

	server := NewServer(mockStore)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/akun/%d", akun.ID)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	requireMatchBody(t, recorder.Body, akun)
}

func randomAkun() database.Akun {
	return database.Akun{
		ID:    util.RandomInt(1, 1000),
		Nama:  util.RandomName(),
		Saldo: util.RandomSaldo(),
	}
}

func requireMatchBody(t *testing.T, body *bytes.Buffer, akun database.Akun) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getAkun database.Akun
	err = json.Unmarshal(data, &getAkun)
	require.NoError(t, err)
	require.Equal(t, akun, getAkun)
}

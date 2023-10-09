package api

import (
	"bytes"
	"database/sql"
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

	testCase := []struct {
		name          string
		id            int64
		buildStubs    func(store *mockdatabase.MockStore)
		checkResponse func(t *testing.T, recorder httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   akun.ID,
			buildStubs: func(store *mockdatabase.MockStore) {
				store.EXPECT().
					GetAkun(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(akun, nil)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireMatchBody(t, recorder.Body, akun)
			},
		},
		// TODO: make new case
		{
			name: "Not Found",
			id:   akun.ID,
			buildStubs: func(store *mockdatabase.MockStore) {
				store.EXPECT().
					GetAkun(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(database.Akun{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "Internal Error",
			id:   akun.ID,
			buildStubs: func(store *mockdatabase.MockStore) {
				store.EXPECT().
					GetAkun(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(database.Akun{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "Bad Request",
			id:   0,
			buildStubs: func(store *mockdatabase.MockStore) {
				store.EXPECT().
					GetAkun(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := mockdatabase.NewMockStore(ctrl)
			tc.buildStubs(mockStore)

			server := NewServer(mockStore)

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/akun/%d", tc.id)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, *recorder)
		})
	}

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

package checkwx

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaf(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/taf/ABCD", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		var resp response
		resp.Data = append(resp.Data, "some taf code")

		require.NoError(t, json.NewEncoder(w).Encode(resp))
	})

	taf, err := client.Taf("ABCD")
	require.NoError(t, err)
	require.Equal(t, "some taf code", taf)
}

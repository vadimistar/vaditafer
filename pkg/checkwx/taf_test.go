package checkwx

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vadimistar/vaditafer/pkg/taf"
)

func TestTaf(t *testing.T) {
	teardown := setup()
	defer teardown()

	var tests = make([]struct {
		resp   data
		expect *taf.Taf
	}, 0)

	var resp data

	resp.Timestamp.Issued = time.Date(2023, 2, 1, 5, 7, 0, 0, time.UTC).Format(timeLayout)
	resp.Timestamp.From = time.Date(2023, 2, 1, 6, 0, 0, 0, time.UTC).Format(timeLayout)
	resp.Timestamp.To = time.Date(2023, 2, 2, 5, 0, 0, 0, time.UTC).Format(timeLayout)

	for _, test := range tests {
		mux.HandleFunc("/taf/ABCD", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)

			require.NoError(t, json.NewEncoder(w).Encode(test.resp))
		})

		got, err := client.Taf("ABCD")
		require.NoError(t, err)
		require.Equal(t, test.expect, got)
	}
}

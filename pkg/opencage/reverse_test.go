package opencage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReverseGeocode(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/geocode/v1/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		q := r.URL.Query()
		if q.Get(client.ApiKey) != client.ApiKey {
			t.Errorf("invalid api key parameter with len = %d", len(client.ApiKey))
		}
		assert.Equal(t, fmt.Sprintf("%f+%f", 30.0, 50.0), q.Get("q"))

		var resp response
		resp.Results = append(resp.Results, result{
			Formatted: "interesting city",
			Geometry: struct {
				Lat float64 "json:\"lat\""
				Lng float64 "json:\"lng\""
			}{Lat: 29.6, Lng: 50.5},
		})

		require.NoError(t, json.NewEncoder(w).Encode(resp))
	})

	place, err := client.ReverseGeocode(30.0, 50.0)
	require.NoError(t, err)
	assert.Equal(t, "interesting city", place)
}

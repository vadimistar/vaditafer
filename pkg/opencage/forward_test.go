package opencage

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestForwardGeocode(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/geocode/v1/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		q := r.URL.Query()
		if q.Get(client.ApiKey) != client.ApiKey {
			t.Errorf("invalid api key parameter with len = %d", len(client.ApiKey))
		}
		assert.Equal(t, "Narnia", q.Get("q"))

		var resp response
		resp.Results = append(resp.Results, result{
			Formatted: "narnia country",
			Geometry: struct {
				Lat float64 "json:\"lat\""
				Lng float64 "json:\"lng\""
			}{Lat: 20.0, Lng: 50.0},
		})

		require.NoError(t, json.NewEncoder(w).Encode(resp))
	})

	lat, lng, err := client.ForwardGeocode("Narnia")
	require.NoError(t, err)
	assert.Equal(t, 20.0, lat)
	assert.Equal(t, 50.0, lng)
}

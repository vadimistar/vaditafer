package avianation

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClosestAirports(t *testing.T) {
	teardown := setup()
	defer teardown()

	mux.HandleFunc("/adds/dataserver_current/httpparam", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/xml")
		w.WriteHeader(http.StatusOK)

		url := r.URL
		q := url.Query()
		assert.Equal(t, []string{"stations"}, q["dataSource"])
		assert.Equal(t, []string{"5"}, q["radialDistance"])
		assert.Equal(t, []string{"xml"}, q["format"])
		assert.Contains(t, r.URL.String(), "&requestType=retrieve")
		_, coords, ok := strings.Cut(r.URL.String(), ";")
		require.True(t, ok)
		assert.Equal(t, fmt.Sprintf("%f,%f", 80.5, 40.5), coords)

		var resp response
		resp.XMLName.Local = "response"
		resp.Data.Station = append(resp.Data.Station, station{
			StationID: "ABCD",
		}, station{
			StationID: "EFGH",
		})

		respBytes, err := xml.Marshal(resp)
		require.NoError(t, err, "cannot marshal response")

		w.Write(respBytes)
	})

	ids, err := client.ClosestAirports(40.5, 80.5, 5)
	require.NoError(t, err)

	require.Equal(t, []string{"ABCD", "EFGH"}, ids)
}

package avianation

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) ClosestAirports(lat float64, lng float64, radialDistance int) (ids []string, err error) {
	defer func() {
		err = errors.Wrapf(err, "client has an api endpoint: %s", c.ApiEndpoint)
	}()
	c.defaultify()

	if lat < -90.0 || lat > 90.0 {
		return nil, fmt.Errorf("latitude value is outside the limits (must be between -90.0 and 90.0): %f", lat)
	}
	if lng < -180.0 || lng > 180.0 {
		return nil, fmt.Errorf("longitude value is outside the limits (must be between -180.0 and 180.0): %f", lng)
	}
	if radialDistance > 500 || radialDistance <= 0 {
		return nil, fmt.Errorf("radial distance is outside the limits (must be 0 <= x <= 500): %d", radialDistance)
	}

	url, err := url.Parse(c.ApiEndpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "parse api endpoint: %s", c.ApiEndpoint)
	}
	url.Path = "adds/dataserver_current/httpparam"

	q := url.Query()
	q.Set("dataSource", "stations")
	q.Set("requestType", "retrieve")
	q.Set("format", "xml")
	url.RawQuery = q.Encode()

	resp, err := http.Get(fmt.Sprintf("%s&radialDistance=%d;%f,%f", url.String(), radialDistance, lng, lat))
	if err != nil {
		return nil, errors.Wrap(err, "cannot access url")
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	var response response
	if err := xml.Unmarshal(data, &response); err != nil {
		return nil, errors.Wrap(err, "unmarshal response")
	}

	for _, station := range response.Data.Station {
		ids = append(ids, station.StationID)
	}

	return ids, nil
}

type response struct {
	XMLName                   xml.Name `xml:"response"`
	Text                      string   `xml:",chardata"`
	Version                   string   `xml:"version,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	RequestIndex              string   `xml:"request_index"`
	DataSource                struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"data_source"`
	Request struct {
		Text   string `xml:",chardata"`
		Type   string `xml:"type,attr"`
		Status struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"status"`
		StayInformed struct {
			Blog     string `json:"blog"`
			Mastodon string `json:"mastodon"`
		} `json:"stay_informed"`
		Thanks    string `json:"thanks"`
		Timestamp struct {
			CreatedHTTP string `json:"created_http"`
			CreatedUnix int    `json:"created_unix"`
		} `json:"timestamp"`
		TotalResults int `json:"total_results"`
	} `xml:"request"`
	Errors      string `xml:"errors"`
	Warnings    string `xml:"warnings"`
	TimeTakenMs string `xml:"time_taken_ms"`
	Data        struct {
		Text       string    `xml:",chardata"`
		NumResults string    `xml:"num_results,attr"`
		Station    []station `xml:"Station"`
	} `xml:"data"`
}

type station struct {
	Text       string `xml:",chardata"`
	StationID  string `xml:"station_id"`
	Latitude   string `xml:"latitude"`
	Longitude  string `xml:"longitude"`
	ElevationM string `xml:"elevation_m"`
	Site       string `xml:"site"`
	Country    string `xml:"country"`
	SiteType   struct {
		Text  string `xml:",chardata"`
		METAR string `xml:"METAR"`
		TAF   string `xml:"TAF"`
	} `xml:"site_type"`
	WmoID string `xml:"wmo_id"`
}

package opencage

import (
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) ForwardGeocode(query string) (lat float64, lng float64, err error) {
	if query == "" {
		return 0, 0, errors.New("empty query")
	}

	q := make(url.Values)
	q.Set("key", c.ApiKey)
	q.Set("q", query)

	resp, err := c.geocode(q)
	if err != nil {
		return 0, 0, err
	}

	if len(resp.Results) <= 0 {
		return 0, 0, errors.New("no results")
	}

	return resp.Results[0].Geometry.Lat, resp.Results[0].Geometry.Lng, nil
}

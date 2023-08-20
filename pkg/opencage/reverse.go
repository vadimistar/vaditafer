package opencage

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) ReverseGeocode(lat float64, lng float64) (place string, err error) {
	if lat < -90.0 || lat > 90.0 {
		return "", fmt.Errorf("latitude value is outside the limits (must be between -90.0 and 90.0): %f", lat)
	}
	if lng < -180.0 || lng > 180.0 {
		return "", fmt.Errorf("longitude value is outside the limits (must be between -180.0 and 180.0): %f", lng)
	}

	q := make(url.Values)
	q.Set("key", c.ApiKey)
	q.Set("q", fmt.Sprintf("%f+%f", lat, lng))

	resp, err := c.geocode(q)
	if err != nil {
		return "", err
	}

	if len(resp.Results) <= 0 {
		return "", errors.New("no results")
	}

	return resp.Results[0].Formatted, nil
}

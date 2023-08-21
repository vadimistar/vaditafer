package opencage

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) geocode(params url.Values) (resp response, err error) {
	defer func() {
		err = errors.Wrapf(err, "client has an api endpoint: %s", c.ApiEndpoint)
	}()
	c.defaultify()

	url, err := url.Parse(c.ApiEndpoint)
	if err != nil {
		return resp, errors.Wrapf(err, "parse api endpoint: %s", c.ApiEndpoint)
	}
	url.Path = "geocode/v1/json"
	url.RawQuery = params.Encode()

	httpResp, err := http.Get(url.String())
	if err != nil {
		return resp, errors.Wrap(err, "cannot access url")
	}

	defer httpResp.Body.Close()

	if err := json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return resp, errors.Wrap(err, "decode response")
	}

	return resp, nil
}

type response struct {
	Results []result `json:"results"`
}

type result struct {
	Annotations struct {
		Timezone struct {
			Name string `json:"name"`
		} `json:"timezone"`
	} `json:"annotations"`
	Formatted string `json:"formatted"`
	Geometry  struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"geometry"`
}

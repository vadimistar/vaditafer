package checkwx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"github.com/pkg/errors"
)

func (c *Client) Taf(id string) (taf string, err error) {
	defer func() {
		err = errors.Wrapf(err, "client has an endpoint: %s", c.ApiEndpoint)
	}()

	if !idRegex.MatchString(id) {
		return "", fmt.Errorf("expect id to be 4 characters long and consist only of english letters: %s", id)
	}

	url, err := url.Parse(c.ApiEndpoint)
	if err != nil {
		return "", errors.Wrap(err, "decode client endpoint")
	}
	url.Path = fmt.Sprintf("taf/%s", id)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return "", errors.Wrap(err, "cannot create request")
	}

	req.Header.Add("X-API-Key", c.ApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "cannot do request")
	}

	defer resp.Body.Close()

	var response response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", errors.Wrap(err, "decode response")
	}

	if len(response.Data) <= 0 {
		return "", errors.New("no response")
	}

	return response.Data[0], nil
}

var idRegex = regexp.MustCompile("^[A-Z]{4}$")

type response struct {
	Data    []string `json:"data"`
	Results int      `json:"results"`
}

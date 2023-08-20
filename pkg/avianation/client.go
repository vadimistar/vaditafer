package avianation

const base = "https://aviationweather.gov"

type Client struct {
	ApiEndpoint string
}

func (c *Client) defaultify() {
	if c.ApiEndpoint == "" {
		c.ApiEndpoint = base
	}
}

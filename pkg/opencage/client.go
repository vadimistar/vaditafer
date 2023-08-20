package opencage

const base = "https://api.opencagedata.com"

type Client struct {
	ApiEndpoint string
	ApiKey      string
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}

func (c *Client) defaultify() {
	if c.ApiEndpoint == "" {
		c.ApiEndpoint = base
	}
}

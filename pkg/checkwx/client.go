package checkwx

const base = "https://api.checkwx.com"

type Client struct {
	ApiEndpoint string
	ApiKey      string
}

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey, ApiEndpoint: base}
}

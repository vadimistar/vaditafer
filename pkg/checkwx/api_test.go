package checkwx

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
	client *Client
)

func setup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = &Client{
		ApiEndpoint: server.URL,
	}

	return func() {
		server.Close()
	}
}

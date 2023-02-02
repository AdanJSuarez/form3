package internal

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	timeout    = 5 * time.Second
	POST       = "POST"
	urlAccount = "/v1/organisation/accounts"
)

type Connection struct {
	url    string
	client http.Client
}

func NewConnection(method, url string) *Connection {
	client := http.Client{
		Timeout: timeout,
	}

	return &Connection{url: url, client: client}
}

func (c *Connection) Post(body io.Reader) (*http.Response, error) {
	url, err := url.JoinPath(string(c.url), urlAccount)
	if err != nil {
		return nil, err
	}

	req := request(POST, url, body)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func request(method string, url string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil
	}
	// TODO: Add needed header
	r.Header.Add("", "")
	return r
}

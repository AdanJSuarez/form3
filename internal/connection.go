package internal

import (
	"io"
	"net/http"
	"time"
)

const timeout = 5 * time.Second
const POST = "POST"

type URL string

type Connection struct {
	url    URL
	client http.Client
}

func NewConnection(method, url URL) *Connection {
	client := http.Client{
		Timeout: timeout,
	}

	return &Connection{url: url, client: client}
}

func (c *Connection) Post(body io.Reader) (*http.Response, error) {
	req := request(POST, c.url, body)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func request(method string, url URL, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil
	}
	// TODO: Add needed header
	r.Header.Add("", "")
	return r
}

package connection

import (
	"io"
	"net/http"
	"time"
)

const (
	timeout = 10 * time.Second
	GET     = "GET"
	POST    = "POST"
	DELETE  = "DELETE"
)

type Connection struct {
	url    string
	client http.Client
}

func New(url string) *Connection {
	client := http.Client{
		Timeout: timeout,
	}
	return &Connection{url: url, client: client}
}

func (c *Connection) Get(value string) (*http.Response, error) {
	request := c.request(GET, c.url, nil)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Connection) Post(body io.Reader) (*http.Response, error) {
	req := c.request(POST, c.url, body)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Connection) Delete() {
	// TODO: Implement delete
}

// request returns a request by http method and URL.
// Easy to set Header if needed.
func (c *Connection) request(method string, url string, body io.Reader) *http.Request {
	r, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil
	}
	return r
}

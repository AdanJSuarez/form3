package connection

import (
	"io"
	"net/http"
	"net/url"
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

// New returns a Connection pointer initialized.
func New(url string) *Connection {
	client := http.Client{
		Timeout: timeout,
	}
	return &Connection{url: url, client: client}
}

func (c *Connection) Get(value string) (*http.Response, error) {
	url, err := url.JoinPath(c.url, value)
	if err != nil {
		return nil, err
	}
	request := c.request(GET, url, nil)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Connection) Post(body io.Reader) (*http.Response, error) {
	request := c.request(POST, c.url, body)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Connection) Delete(value string) (*http.Response, error) {
	url, err := url.JoinPath(c.url, value)
	if err != nil {
		return nil, err
	}
	request := c.request(DELETE, url, nil)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// request returns a request by http method and URL.
// Easy to set Header if needed.
func (c *Connection) request(method string, url string, body io.Reader) *http.Request {
	request, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil
	}
	return request
}

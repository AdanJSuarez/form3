package client

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

type Client struct {
	url    string
	client http.Client
}

func New(url string) *Client {
	client := http.Client{
		Timeout: timeout,
	}
	return &Client{url: url, client: client}
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := url.JoinPath(c.url, value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(GET, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Post(body io.Reader) (*http.Response, error) {
	request, err := c.request(POST, c.url, body)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Delete(value string) (*http.Response, error) {
	url, err := url.JoinPath(c.url, value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(DELETE, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// request returns a request by http method and URL.
// Easy to set Header if needed.
func (c *Client) request(method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, string(url), body)
	if err != nil {
		return nil, err
	}
	return request, nil
}

package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	timeout     = 10 * time.Second
	GET         = "GET"
	POST        = "POST"
	DELETE      = "DELETE"
	emptySize   = 0
	defaultHost = "api.form3.tech"
)

type Client struct {
	url    string
	client httpClient
}

func New(url string) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{url: url, client: client}
}

func (c *Client) Get(value string) (*http.Response, error) {
	url, err := c.joinValueToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(GET, url, NewRequestBody(nil, emptySize))
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Post(body *RequestBody) (*http.Response, error) {
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
	url, err := c.joinValueToURL(value)
	if err != nil {
		return nil, err
	}

	request, err := c.request(DELETE, url, NewRequestBody(nil, emptySize))
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
func (c *Client) request(method string, url string, body *RequestBody) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body.Body())
	if err != nil {
		return nil, err
	}

	c.addRequiredHeader(request)

	if body.Body() != nil {
		c.addHeaderToRequestWithBody(request, body.Size())
	}

	return request, nil
}

func (c *Client) joinValueToURL(value string) (string, error) {
	url, err := url.JoinPath(c.url, value)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (c *Client) addRequiredHeader(request *http.Request) {
	request.Header.Add("Host", c.getHostFromURL())
	request.Header.Add("Date", time.Now().String())
	request.Header.Add("Accept", "application/vnd.api+json")
}

func (c *Client) addHeaderToRequestWithBody(request *http.Request, size int) {
	request.Header.Add("Content-Type", "application/vnd.api+json")
	request.Header.Add("Content-Length", fmt.Sprint(size))
}

func (c *Client) getHostFromURL() string {
	u, err := url.Parse(c.url)
	if err != nil {
		return defaultHost
	}
	return u.Host
}
